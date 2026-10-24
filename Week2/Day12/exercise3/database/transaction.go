package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
)

const (
	maxTransferAttempts = 3
	initialRetryDelay   = 50 * time.Millisecond
)

func TransferBalance(
	ctx context.Context,
	db *sql.DB,
	fromID int,
	toID int,
	amount int64,
) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	if fromID == toID {
		return errors.New("source and destination must be different")
	}

	if err := ctx.Err(); err != nil {
		return fmt.Errorf("transaction context canceled: %w", err)
	}

	return transferBalanceWithRetry(
		ctx,
		func() error {
			return transferBalanceOnce(ctx, db, fromID, toID, amount)
		},
		isDeadlockError,
	)
}

// transferBalanceWithRetry retries only errors identified as deadlocks.
// Each attempt must start a completely new transaction.
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
				attempt,
				err,
			)
		}

		delay := initialRetryDelay * time.Duration(1<<(attempt-1))
		timer := time.NewTimer(delay)

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
	fromID int,
	toID int,
	amount int64,
) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Lock accounts in a consistent ID order to reduce deadlock risk.
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

	result, err := tx.ExecContext(
		ctx,
		`UPDATE accounts
		 SET balance = balance - $1
		 WHERE id = $2 AND balance >= $1`,
		amount,
		fromID,
	)
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

	result, err = tx.ExecContext(
		ctx,
		`UPDATE accounts
		 SET balance = balance + $1
		 WHERE id = $2`,
		amount,
		toID,
	)
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

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
