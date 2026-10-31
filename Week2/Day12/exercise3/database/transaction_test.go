package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
)

func TestTransferBalance_InvalidAmount(t *testing.T) {
	err := TransferBalance(context.Background(), nil, 1, 2, 0)

	if err == nil {
		t.Fatal("expected error for zero amount")
	}

	if !strings.Contains(err.Error(), "positive") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTransferBalance_NegativeAmount(t *testing.T) {
	err := TransferBalance(context.Background(), nil, 1, 2, -100)

	if err == nil {
		t.Fatal("expected error for negative amount")
	}
}

func TestTransferBalance_SameAccount(t *testing.T) {
	err := TransferBalance(context.Background(), nil, 1, 1, 100)

	if err == nil {
		t.Fatal("expected error when source and destination are the same")
	}

	if !strings.Contains(err.Error(), "different") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTransferBalance_CanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := TransferBalance(ctx, nil, 1, 2, 100)

	if err == nil {
		t.Fatal("expected error for canceled context")
	}

	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got: %v", err)
	}
}

func TestIsDeadlockError_DirectDeadlock(t *testing.T) {
	err := &pq.Error{Code: "40P01"}

	if !isDeadlockError(err) {
		t.Fatal("expected PostgreSQL deadlock error to be detected")
	}
}

func TestIsDeadlockError_WrappedDeadlock(t *testing.T) {
	err := fmt.Errorf(
		"transfer failed: %w",
		&pq.Error{Code: "40P01"},
	)

	if !isDeadlockError(err) {
		t.Fatal("expected wrapped PostgreSQL deadlock error to be detected")
	}
}

func TestIsDeadlockError_OtherSQLState(t *testing.T) {
	err := &pq.Error{Code: "23505"} // Unique constraint violation

	if isDeadlockError(err) {
		t.Fatal("expected non-deadlock error not to be detected as a deadlock")
	}
}

func TestIsDeadlockError_OrdinaryError(t *testing.T) {
	err := errors.New("database unavailable")

	if isDeadlockError(err) {
		t.Fatal("expected ordinary error not to be detected as a deadlock")
	}
}

func TestTransferBalanceWithRetry_SucceedsAfterDeadlock(t *testing.T) {
	ctx := context.Background()
	attempts := 0

	err := transferBalanceWithRetry(
		ctx,
		func() error {
			attempts++
			if attempts < 3 {
				return &pq.Error{Code: "40P01"}
			}
			return nil
		},
		isDeadlockError,
	)

	if err != nil {
		t.Fatalf("expected retry to succeed, got: %v", err)
	}

	if attempts != 3 {
		t.Fatalf("expected 3 attempts, got %d", attempts)
	}
}

func TestTransferBalanceWithRetry_NoRetryForOtherError(t *testing.T) {
	attempts := 0
	expectedErr := errors.New("database unavailable")

	err := transferBalanceWithRetry(
		context.Background(),
		func() error {
			attempts++
			return expectedErr
		},
		isDeadlockError,
	)

	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected original error, got: %v", err)
	}

	if attempts != 1 {
		t.Fatalf("expected 1 attempt, got %d", attempts)
	}
}

func TestTransferBalanceWithRetry_StopsAfterMaxAttempts(t *testing.T) {
	attempts := 0

	err := transferBalanceWithRetry(
		context.Background(),
		func() error {
			attempts++
			return &pq.Error{Code: "40P01"}
		},
		isDeadlockError,
	)

	if err == nil {
		t.Fatal("expected error after exhausting retries")
	}

	if !strings.Contains(err.Error(), "after 3 attempts") {
		t.Errorf("expected retry exhaustion message, got: %v", err)
	}

	if !errors.As(err, new(*pq.Error)) {
		t.Errorf("expected the original PostgreSQL error to remain accessible, got: %v", err)
	}

	if attempts != maxTransferAttempts {
		t.Fatalf(
			"expected %d attempts, got %d",
			maxTransferAttempts,
			attempts,
		)
	}
}

func TestTransferBalanceWithRetry_CanceledDuringRetry(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	attempts := 0

	err := transferBalanceWithRetry(
		ctx,
		func() error {
			attempts++
			cancel()

			return &pq.Error{Code: "40P01"}
		},
		isDeadlockError,
	)

	if err == nil {
		t.Fatal("expected error when context is canceled during retry")
	}

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got: %v", err)
	}

	if attempts != 1 {
		t.Fatalf("expected 1 attempt, got %d", attempts)
	}
}

func TestTransferBalanceWithRetry_RollbackOnCreditFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	mock.ExpectBegin()

	// Lock the source account first.
	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Lock the destination account.
	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	// Debit succeeds.
	mock.ExpectExec("UPDATE accounts").
		WithArgs(int64(100), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Credit fails.
	mock.ExpectExec("UPDATE accounts").
		WithArgs(int64(100), 2).
		WillReturnError(errors.New("simulated credit failure"))

	// The transaction must roll back.
	mock.ExpectRollback()

	err = TransferBalance(ctx, db, 1, 2, 100)
	if err == nil {
		t.Fatal("expected transfer to fail when credit operation fails")
	}

	if !strings.Contains(err.Error(), "failed to credit account") {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}
func TestTransferBalanceWithKey_EmptyKey(t *testing.T) {
	err := TransferBalanceWithKey(
		context.Background(),
		nil,
		"   ",
		1,
		2,
		100,
	)

	if err == nil {
		t.Fatal("expected error for empty idempotency key")
	}

	if err.Error() != "idempotency key must not be empty" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTransferBalanceWithKey_InvalidAmount(t *testing.T) {
	err := TransferBalanceWithKey(
		context.Background(),
		nil,
		"test-key",
		1,
		2,
		0,
	)

	if err == nil {
		t.Fatal("expected error for invalid amount")
	}

	if err.Error() != "transfer amount must be positive" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTransferBalanceWithKey_SameAccount(t *testing.T) {
	err := TransferBalanceWithKey(
		context.Background(),
		nil,
		"test-key",
		1,
		1,
		100,
	)

	if err == nil {
		t.Fatal("expected error for same source and destination")
	}

	if err.Error() != "source and destination must be different" {
		t.Errorf("unexpected error: %v", err)
	}
}
func TestTransferBalanceWithKey_CanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := TransferBalanceWithKey(
		ctx,
		nil,
		"test-key",
		1,
		2,
		100,
	)

	if err == nil {
		t.Fatal("expected error for canceled context")
	}

	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}

func TestTransferBalanceWithRetry_ContextCanceledBeforeAttempt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	attempts := 0

	err := transferBalanceWithRetry(
		ctx,
		func() error {
			attempts++
			return nil
		},
		isDeadlockError,
	)

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got: %v", err)
	}

	if attempts != 0 {
		t.Fatalf("expected 0 attempts, got %d", attempts)
	}
}

func TestTransferBalanceWithKeyOnce_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("unit-test-key", 1, 2, int64(100)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// transferBalanceTx locks both accounts in ascending ID order.
	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	mock.ExpectExec("UPDATE accounts").
		WithArgs(int64(100), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec("UPDATE accounts").
		WithArgs(int64(100), 2).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	err = transferBalanceWithKeyOnce(ctx, db, "unit-test-key", 1, 2, 100)
	if err != nil {
		t.Fatalf("expected successful transfer, got: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_DuplicateRequest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	mock.ExpectBegin()

	// The idempotency key already exists, so no new row is inserted.
	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("existing-key", 1, 2, int64(100)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	mock.ExpectQuery("SELECT from_account_id, to_account_id, amount").
		WithArgs("existing-key").
		WillReturnRows(
			sqlmock.NewRows([]string{
				"from_account_id",
				"to_account_id",
				"amount",
			}).AddRow(1, 2, int64(100)),
		)

	// The duplicate request commits without transferring money again.
	mock.ExpectCommit()

	err = transferBalanceWithKeyOnce(
		ctx, db, "existing-key", 1, 2, 100,
	)
	if err != nil {
		t.Fatalf("expected duplicate request to succeed, got: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_IdempotencyKeyConflict(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("existing-key", 1, 2, int64(200)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	mock.ExpectQuery("SELECT from_account_id, to_account_id, amount").
		WithArgs("existing-key").
		WillReturnRows(
			sqlmock.NewRows([]string{
				"from_account_id",
				"to_account_id",
				"amount",
			}).AddRow(1, 2, int64(100)),
		)

	err = transferBalanceWithKeyOnce(
		context.Background(), db, "existing-key", 1, 2, 200,
	)

	if !errors.Is(err, ErrIdempotencyKeyConflict) {
		t.Fatalf(
			"expected ErrIdempotencyKeyConflict, got: %v",
			err,
		)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_InsertRequestError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("insert-error-key", 1, 2, int64(100)).
		WillReturnError(errors.New("database insert failed"))

	err = transferBalanceWithKeyOnce(
		context.Background(), db, "insert-error-key", 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected an error when inserting transfer request fails")
	}

	if !strings.Contains(err.Error(), "failed to record transfer request") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_RowsAffectedError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("rows-error-key", 1, 2, int64(100)).
		WillReturnResult(sqlmock.NewErrorResult(
			errors.New("rows affected unavailable"),
		))

	err = transferBalanceWithKeyOnce(
		context.Background(), db, "rows-error-key", 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected RowsAffected error")
	}

	if !strings.Contains(err.Error(), "failed to check transfer request") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_ReadExistingRequestError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("read-error-key", 1, 2, int64(100)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	mock.ExpectQuery("SELECT from_account_id, to_account_id, amount").
		WithArgs("read-error-key").
		WillReturnError(errors.New("query failed"))

	err = transferBalanceWithKeyOnce(
		context.Background(), db, "read-error-key", 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected an error when reading existing transfer request")
	}

	if !strings.Contains(err.Error(), "failed to read existing transfer request") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_DuplicateCommitError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("duplicate-commit-key", 1, 2, int64(100)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	rows := sqlmock.NewRows([]string{
		"from_account_id",
		"to_account_id",
		"amount",
	}).AddRow(1, 2, int64(100))

	mock.ExpectQuery("SELECT from_account_id, to_account_id, amount").
		WithArgs("duplicate-commit-key").
		WillReturnRows(rows)

	mock.ExpectCommit().
		WillReturnError(errors.New("commit failed"))

	err = transferBalanceWithKeyOnce(
		context.Background(), db, "duplicate-commit-key", 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected duplicate request commit error")
	}

	if !strings.Contains(err.Error(), "failed to commit duplicate request") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_BeginTransactionError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin().
		WillReturnError(errors.New("begin transaction failed"))

	err = transferBalanceWithKeyOnce(
		context.Background(), db, "begin-error-key", 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected transaction begin error")
	}

	if !strings.Contains(err.Error(), "failed to begin transaction") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_RollbackOnLockFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("lock-error-key", 1, 2, int64(100)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Account 1 lock succeeds.
	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Account 2 lock fails.
	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(2).
		WillReturnError(errors.New("database lock error"))

	mock.ExpectRollback()

	err = transferBalanceWithKeyOnce(
		context.Background(), db, "lock-error-key", 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected account lock error")
	}

	if !strings.Contains(err.Error(), "failed to lock account 2") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_RollbackOnDebitFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("debit-error-key", 1, 2, int64(100)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	mock.ExpectExec("UPDATE accounts").
		WithArgs(int64(100), 1).
		WillReturnError(errors.New("debit failed"))

	mock.ExpectRollback()

	err = transferBalanceWithKeyOnce(
		context.Background(), db, "debit-error-key", 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected debit error")
	}

	if !strings.Contains(err.Error(), "failed to debit account") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_RollbackOnCreditFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("credit-error-key", 1, 2, int64(100)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Lock both accounts successfully.
	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	// Debit succeeds.
	mock.ExpectExec("UPDATE accounts").
		WithArgs(int64(100), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Credit fails.
	mock.ExpectExec("UPDATE accounts").
		WithArgs(int64(100), 2).
		WillReturnError(errors.New("credit failed"))

	mock.ExpectRollback()

	err = transferBalanceWithKeyOnce(
		context.Background(), db, "credit-error-key", 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected credit error")
	}

	if !strings.Contains(err.Error(), "failed to credit account") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceWithKeyOnce_TransferCommitError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO transfer_requests").
		WithArgs("transfer-commit-error-key", 1, 2, int64(100)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Lock both accounts.
	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	// Debit succeeds.
	mock.ExpectExec("UPDATE accounts").
		WithArgs(int64(100), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Credit succeeds.
	mock.ExpectExec("UPDATE accounts").
		WithArgs(int64(100), 2).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Final commit fails.
	mock.ExpectCommit().
		WillReturnError(errors.New("transfer commit failed"))

	err = transferBalanceWithKeyOnce(
		context.Background(),
		db,
		"transfer-commit-error-key",
		1,
		2,
		100,
	)

	if err == nil {
		t.Fatal("expected transfer commit error")
	}

	if !strings.Contains(err.Error(), "failed to commit transfer") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}

func TestTransferBalanceTx_AccountNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectRollback()

	err = transferBalanceTx(
		context.Background(), tx, 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected account not found error")
	}

	if !strings.Contains(err.Error(), "account 1 not found") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := tx.Rollback(); err != nil {
		t.Errorf("failed to rollback transaction: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}
func TestTransferBalanceTx_DebitRowsAffectedError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	mock.ExpectExec("UPDATE accounts SET balance = balance -").
		WithArgs(int64(100), 1).
		WillReturnResult(sqlmock.NewErrorResult(errors.New("rows affected failed")))

	mock.ExpectRollback()

	err = transferBalanceTx(
		context.Background(), tx, 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected debit rows affected error")
	}

	if !strings.Contains(err.Error(), "failed to check source account") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := tx.Rollback(); err != nil {
		t.Errorf("failed to rollback transaction: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}
func TestTransferBalanceTx_InsufficientBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	mock.ExpectExec("UPDATE accounts SET balance = balance -").
		WithArgs(int64(100), 1).
		WillReturnResult(sqlmock.NewResult(0, 0))

	mock.ExpectRollback()

	err = transferBalanceTx(
		context.Background(), tx, 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected insufficient balance error")
	}

	if !strings.Contains(err.Error(), "insufficient balance") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := tx.Rollback(); err != nil {
		t.Errorf("failed to rollback transaction: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}
func TestTransferBalanceTx_CreditRowsAffectedError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	mock.ExpectExec("UPDATE accounts SET balance = balance -").
		WithArgs(int64(100), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec("UPDATE accounts SET balance = balance \\+").
		WithArgs(int64(100), 2).
		WillReturnResult(sqlmock.NewErrorResult(errors.New("rows affected failed")))

	mock.ExpectRollback()

	err = transferBalanceTx(
		context.Background(), tx, 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected credit rows affected error")
	}

	if !strings.Contains(err.Error(), "failed to check destination account") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := tx.Rollback(); err != nil {
		t.Errorf("failed to rollback transaction: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}
func TestTransferBalanceTx_DestinationNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery("SELECT id FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	mock.ExpectExec("UPDATE accounts SET balance = balance -").
		WithArgs(int64(100), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec("UPDATE accounts SET balance = balance \\+").
		WithArgs(int64(100), 2).
		WillReturnResult(sqlmock.NewResult(0, 0))

	mock.ExpectRollback()

	err = transferBalanceTx(
		context.Background(), tx, 1, 2, 100,
	)

	if err == nil {
		t.Fatal("expected destination account not found error")
	}

	if !strings.Contains(err.Error(), "destination account not found") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if err := tx.Rollback(); err != nil {
		t.Errorf("failed to rollback transaction: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expectations: %v", err)
	}
}
