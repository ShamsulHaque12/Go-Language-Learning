package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

const (
	maxTransferAttempts = 3
	initialRetryDelay   = 50 * time.Millisecond
)

var ErrIdempotencyKeyConflict = errors.New(
	"idempotency key already used for a different transfer",
)

func TransferBalance(
	ctx context.Context,
	db *sql.DB,
	fromID, toID int,
	amount int64,
) error {
	if err := validateTransfer(fromID, toID, amount); err != nil {
		return err
	}

	if err := ctx.Err(); err != nil {
		return fmt.Errorf("transaction context canceled: %w", err)
	}

	return transferBalanceWithRetry(ctx, func() error {
		return transferBalanceOnce(ctx, db, fromID, toID, amount)
	}, isDeadlockError)
}

func TransferBalanceWithKey(
	ctx context.Context,
	db *sql.DB,
	key string,
	fromID, toID int,
	amount int64,
) error {
	if strings.TrimSpace(key) == "" {
		return errors.New("idempotency key must not be empty")
	}

	if err := validateTransfer(fromID, toID, amount); err != nil {
		return err
	}

	if err := ctx.Err(); err != nil {
		return fmt.Errorf("transaction context canceled: %w", err)
	}

	return transferBalanceWithRetry(ctx, func() error {
		return transferBalanceWithKeyOnce(
			ctx, db, key, fromID, toID, amount,
		)
	}, isDeadlockError)
}

func validateTransfer(fromID, toID int, amount int64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}
	if fromID == toID {
		return errors.New("source and destination must be different")
	}
	return nil
}

// Every retry must start a fresh SQL transaction.
func transferBalanceWithRetry(
	ctx context.Context,
	transfer func() error,
	isDeadlock func(error) bool,
) error {
	for attempt := 1; attempt <= maxTransferAttempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("transaction context canceled: %w", err)
		}

		err := transfer()
		if err == nil {
			return nil
		}

		if !isDeadlock(err) {
			return err
		}

		if attempt == maxTransferAttempts {
			return fmt.Errorf(
				"transfer failed after %d attempts due to deadlock: %w",
				attempt, err,
			)
		}

		timer := time.NewTimer(
			initialRetryDelay * time.Duration(1<<(attempt-1)),
		)

		select {
		case <-ctx.Done():
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			return fmt.Errorf(
				"transaction context canceled during retry: %w",
				ctx.Err(),
			)
		case <-timer.C:
		}
	}

	return nil
}

func isDeadlockError(err error) bool {
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.SQLState() == "40P01"
}

func transferBalanceOnce(
	ctx context.Context,
	db *sql.DB,
	fromID, toID int,
	amount int64,
) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := transferBalanceTx(ctx, tx, fromID, toID, amount); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func transferBalanceWithKeyOnce(
	ctx context.Context,
	db *sql.DB,
	key string,
	fromID, toID int,
	amount int64,
) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx, `
		INSERT INTO transfer_requests
			(idempotency_key, from_account_id, to_account_id, amount)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (idempotency_key) DO NOTHING
	`, key, fromID, toID, amount)
	if err != nil {
		return fmt.Errorf("failed to record transfer request: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check transfer request: %w", err)
	}

	if rows == 0 {
		var existingFrom, existingTo int
		var existingAmount int64

		err := tx.QueryRowContext(ctx, `
			SELECT from_account_id, to_account_id, amount
			FROM transfer_requests
			WHERE idempotency_key = $1
		`, key).Scan(&existingFrom, &existingTo, &existingAmount)
		if err != nil {
			return fmt.Errorf("failed to read existing transfer request: %w", err)
		}

		if existingFrom != fromID ||
			existingTo != toID ||
			existingAmount != amount {
			return ErrIdempotencyKeyConflict
		}

		// This key already completed the same transfer.
		// Commit the no-op transaction and do not move money again.
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit duplicate request: %w", err)
		}
		return nil
	}

	if err := transferBalanceTx(ctx, tx, fromID, toID, amount); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transfer: %w", err)
	}
	return nil
}

// Locks accounts in ascending ID order to reduce deadlock risk.
// The caller owns the transaction and is responsible for committing it.
func transferBalanceTx(
	ctx context.Context,
	tx *sql.Tx,
	fromID, toID int,
	amount int64,
) error {
	firstID, secondID := fromID, toID
	if firstID > secondID {
		firstID, secondID = secondID, firstID
	}

	for _, accountID := range []int{firstID, secondID} {
		var lockedID int
		err := tx.QueryRowContext(
			ctx,
			`SELECT id FROM accounts WHERE id = $1 FOR UPDATE`,
			accountID,
		).Scan(&lockedID)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("account %d not found", accountID)
			}
			return fmt.Errorf("failed to lock account %d: %w", accountID, err)
		}
	}

	result, err := tx.ExecContext(ctx, `
		UPDATE accounts
		SET balance = balance - $1
		WHERE id = $2 AND balance >= $1
	`, amount, fromID)
	if err != nil {
		return fmt.Errorf("failed to debit account: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check source account: %w", err)
	}
	if rows != 1 {
		return errors.New("insufficient balance")
	}

	result, err = tx.ExecContext(ctx, `
		UPDATE accounts
		SET balance = balance + $1
		WHERE id = $2
	`, amount, toID)
	if err != nil {
		return fmt.Errorf("failed to credit account: %w", err)
	}

	rows, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check destination account: %w", err)
	}
	if rows != 1 {
		return errors.New("destination account not found")
	}

	return nil
}
