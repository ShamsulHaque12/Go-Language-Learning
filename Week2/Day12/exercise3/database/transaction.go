package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// সব transfer-এ account ID-এর একই ক্রমে lock নেওয়া।
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

	// Sender-এর পর্যাপ্ত balance আছে কি না যাচাই করে debit করা।
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

	// Receiver-এর balance বাড়ানো।
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
