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

	// টাকা পাঠানোর account থেকে balance কমানো।
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
		return errors.New("source account not found or insufficient balance")
	}

	// টাকা গ্রহণকারী account-এ balance যোগ করা।
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

	// সব operation সফল হলে পরিবর্তন স্থায়ী করা।
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
