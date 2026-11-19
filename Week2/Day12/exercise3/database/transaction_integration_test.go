package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func setupAccountsTest(t *testing.T) (*sql.DB, func(), string) {
	t.Helper()

	// Prefer an explicitly exported environment variable. If it is absent,
	// try the exercise's .env file without printing credentials.
	if os.Getenv("DATABASE_URL") == "" {
		for _, envPath := range []string{
			"../.env",
			"Week2/Day12/exercise3/.env",
			".env",
		} {
			if err := godotenv.Load(envPath); err == nil && os.Getenv("DATABASE_URL") != "" {
				break
			}
		}
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("Skipping integration test: DATABASE_URL is not set")
	}

	// A unique prefix ensures cleanup only touches records created by this run.
	testPrefix := fmt.Sprintf("GoTest_%d_", time.Now().UnixNano())

	db, err := ConnectDB(databaseURL)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			balance BIGINT NOT NULL CHECK (balance >= 0)
		)
	`)
	if err != nil {
		db.Close()
		t.Fatalf("failed to create accounts table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS transfer_requests (
			idempotency_key TEXT PRIMARY KEY,
			from_account_id INTEGER NOT NULL,
			to_account_id INTEGER NOT NULL,
			amount BIGINT NOT NULL CHECK (amount > 0),
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		db.Close()
		t.Fatalf("failed to create transfer_requests table: %v", err)
	}

	cleanupTestData := func() error {
		// Use a literal prefix comparison. LIKE would treat underscores as
		// wildcards, which could match unrelated records.
		_, err := db.Exec(`
			DELETE FROM transfer_requests
			WHERE LEFT(idempotency_key, LENGTH($1)) = $1
		`, testPrefix)
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			DELETE FROM accounts
			WHERE LEFT(name, LENGTH($1)) = $1
		`, testPrefix)
		return err
	}

	// Never clear old test data before a test starts. Only clean this run's
	// uniquely prefixed rows after the test has completed.
	cleanup := func() {
		if err := cleanupTestData(); err != nil {
			t.Errorf("failed to clean this test's data: %v", err)
		}
		if err := db.Close(); err != nil {
			t.Errorf("failed to close database: %v", err)
		}
	}

	return db, cleanup, testPrefix
}

func createTestAccounts(
	t *testing.T,
	db *sql.DB,
	testPrefix string,
	senderBalance, receiverBalance int64,
) (int, int) {
	t.Helper()

	var fromID, toID int

	err := db.QueryRow(`
		INSERT INTO accounts (name, balance)
		VALUES ($1 || 'Sender', $2)
		RETURNING id
	`, testPrefix, senderBalance).Scan(&fromID)
	if err != nil {
		t.Fatal(err)
	}

	err = db.QueryRow(`
		INSERT INTO accounts (name, balance)
		VALUES ($1 || 'Receiver', $2)
		RETURNING id
	`, testPrefix, receiverBalance).Scan(&toID)
	if err != nil {
		t.Fatal(err)
	}

	return fromID, toID
}

func getBalance(t *testing.T, db *sql.DB, accountID int) int64 {
	t.Helper()

	var balance int64
	err := db.QueryRow(
		`SELECT balance FROM accounts WHERE id = $1`,
		accountID,
	).Scan(&balance)
	if err != nil {
		t.Fatal(err)
	}

	return balance
}

func TestTransferBalance_Commit_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	fromID, toID := createTestAccounts(t, db, testPrefix, 1000, 200)

	if err := TransferBalance(
		context.Background(), db, fromID, toID, 300,
	); err != nil {
		t.Fatalf("transfer failed: %v", err)
	}

	if got := getBalance(t, db, fromID); got != 700 {
		t.Errorf("expected sender balance 700, got %d", got)
	}
	if got := getBalance(t, db, toID); got != 500 {
		t.Errorf("expected receiver balance 500, got %d", got)
	}
}

func TestTransferBalance_Rollback_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	fromID, _ := createTestAccounts(t, db, testPrefix, 1000, 200)

	// Use a positive account ID that does not exist.
	var missingID int
	err := db.QueryRow(`
		SELECT COALESCE(MAX(id), 0) + 1000000
		FROM accounts
	`).Scan(&missingID)
	if err != nil {
		t.Fatalf("failed to generate missing account ID: %v", err)
	}

	err = TransferBalance(
		context.Background(), db, fromID, missingID, 300,
	)
	if err == nil {
		t.Fatal("expected transfer to fail for nonexistent destination")
	}

	if got := getBalance(t, db, fromID); got != 1000 {
		t.Errorf("expected sender balance 1000 after failure, got %d", got)
	}
}

func TestTransferBalance_Concurrent_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	fromID, toID := createTestAccounts(t, db, testPrefix, 10000, 0)

	const (
		workers = 10
		amount  = int64(100)
	)

	var wg sync.WaitGroup
	errs := make(chan error, workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs <- TransferBalance(
				context.Background(), db, fromID, toID, amount,
			)
		}()
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			t.Errorf("concurrent transfer failed: %v", err)
		}
	}

	if got := getBalance(t, db, fromID); got != 9000 {
		t.Errorf("expected sender balance 9000, got %d", got)
	}
	if got := getBalance(t, db, toID); got != 1000 {
		t.Errorf("expected receiver balance 1000, got %d", got)
	}
}

func TestTransferBalance_OppositeDirections_Concurrent_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	accountA, accountB := createTestAccounts(t, db, testPrefix, 10000, 10000)

	const workers = 20

	start := make(chan struct{})
	errs := make(chan error, workers)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		fromID, toID := accountA, accountB
		if i%2 != 0 {
			fromID, toID = accountB, accountA
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			errs <- TransferBalance(
				context.Background(), db, fromID, toID, 100,
			)
		}()
	}

	close(start)
	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			t.Errorf("opposite-direction transfer failed: %v", err)
		}
	}

	if got := getBalance(t, db, accountA); got != 10000 {
		t.Errorf("expected account A balance 10000, got %d", got)
	}
	if got := getBalance(t, db, accountB); got != 10000 {
		t.Errorf("expected account B balance 10000, got %d", got)
	}
}

func TestTransferBalanceWithKey_SameRequestTwice_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	fromID, toID := createTestAccounts(t, db, testPrefix, 1000, 200)
	ctx := context.Background()
	key := testPrefix + "same-request"

	for i := 0; i < 2; i++ {
		if err := TransferBalanceWithKey(
			ctx, db, key, fromID, toID, 300,
		); err != nil {
			t.Fatalf("attempt %d failed: %v", i+1, err)
		}
	}

	if got := getBalance(t, db, fromID); got != 700 {
		t.Errorf("expected sender balance 700, got %d", got)
	}
	if got := getBalance(t, db, toID); got != 500 {
		t.Errorf("expected receiver balance 500, got %d", got)
	}
}

func TestTransferBalanceWithKey_DifferentAmountConflict_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	fromID, toID := createTestAccounts(t, db, testPrefix, 1000, 200)
	ctx := context.Background()
	key := testPrefix + "amount-conflict"

	if err := TransferBalanceWithKey(
		ctx, db, key, fromID, toID, 300,
	); err != nil {
		t.Fatal(err)
	}

	err := TransferBalanceWithKey(
		ctx, db, key, fromID, toID, 400,
	)
	if !errors.Is(err, ErrIdempotencyKeyConflict) {
		t.Fatalf("expected ErrIdempotencyKeyConflict, got %v", err)
	}

	if got := getBalance(t, db, fromID); got != 700 {
		t.Errorf("expected sender balance 700, got %d", got)
	}
	if got := getBalance(t, db, toID); got != 500 {
		t.Errorf("expected receiver balance 500, got %d", got)
	}
}

func TestTransferBalanceWithKey_DifferentAccountsConflict_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	fromID, toID := createTestAccounts(t, db, testPrefix, 1000, 200)
	otherFromID, otherToID := createTestAccounts(t, db, testPrefix+"other_", 800, 100)
	ctx := context.Background()
	key := testPrefix + "account-conflict"

	if err := TransferBalanceWithKey(
		ctx, db, key, fromID, toID, 100,
	); err != nil {
		t.Fatal(err)
	}

	err := TransferBalanceWithKey(
		ctx, db, key, otherFromID, otherToID, 100,
	)
	if !errors.Is(err, ErrIdempotencyKeyConflict) {
		t.Fatalf("expected ErrIdempotencyKeyConflict, got %v", err)
	}

	if got := getBalance(t, db, otherFromID); got != 800 {
		t.Errorf("other sender should remain at 800, got %d", got)
	}
	if got := getBalance(t, db, otherToID); got != 100 {
		t.Errorf("other receiver should remain at 100, got %d", got)
	}
}

func TestTransferBalanceWithKey_DifferentKeys_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	fromID, toID := createTestAccounts(t, db, testPrefix, 1000, 200)
	ctx := context.Background()

	for _, key := range []string{
		testPrefix + "key-one",
		testPrefix + "key-two",
	} {
		if err := TransferBalanceWithKey(
			ctx, db, key, fromID, toID, 100,
		); err != nil {
			t.Fatalf("transfer with key %q failed: %v", key, err)
		}
	}

	if got := getBalance(t, db, fromID); got != 800 {
		t.Errorf("expected sender balance 800, got %d", got)
	}
	if got := getBalance(t, db, toID); got != 400 {
		t.Errorf("expected receiver balance 400, got %d", got)
	}
}

func TestTransferBalanceWithKey_FailedTransferCanRetry_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	fromID, toID := createTestAccounts(t, db, testPrefix, 100, 200)
	ctx := context.Background()
	key := testPrefix + "retry-after-failure"

	err := TransferBalanceWithKey(ctx, db, key, fromID, toID, 300)
	if err == nil {
		t.Fatal("expected insufficient balance error")
	}

	if got := getBalance(t, db, fromID); got != 100 {
		t.Errorf("sender balance changed after failed transfer: %d", got)
	}
	if got := getBalance(t, db, toID); got != 200 {
		t.Errorf("receiver balance changed after failed transfer: %d", got)
	}

	// Correct the balance, then retry with the same key.
	_, err = db.Exec(
		`UPDATE accounts SET balance = 500 WHERE id = $1`,
		fromID,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := TransferBalanceWithKey(
		ctx, db, key, fromID, toID, 300,
	); err != nil {
		t.Fatalf("retry after failed transfer failed: %v", err)
	}

	if got := getBalance(t, db, fromID); got != 200 {
		t.Errorf("expected sender balance 200, got %d", got)
	}
	if got := getBalance(t, db, toID); got != 500 {
		t.Errorf("expected receiver balance 500, got %d", got)
	}
}

func TestTransferBalanceWithKey_ConcurrentSameKey_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	fromID, toID := createTestAccounts(t, db, testPrefix, 1000, 200)
	ctx := context.Background()
	key := testPrefix + "concurrent-same-key"

	const workers = 10

	start := make(chan struct{})
	errs := make(chan error, workers)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			errs <- TransferBalanceWithKey(
				ctx, db, key, fromID, toID, 300,
			)
		}()
	}

	close(start)
	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			t.Errorf("concurrent idempotent transfer failed: %v", err)
		}
	}

	if got := getBalance(t, db, fromID); got != 700 {
		t.Errorf("expected sender balance 700, got %d", got)
	}
	if got := getBalance(t, db, toID); got != 500 {
		t.Errorf("expected receiver balance 500, got %d", got)
	}
}

func TestTransferBalanceWithKey_EmptyKey_Integration(t *testing.T) {
	db, cleanup, testPrefix := setupAccountsTest(t)
	defer cleanup()

	fromID, toID := createTestAccounts(t, db, testPrefix, 1000, 200)

	err := TransferBalanceWithKey(
		context.Background(), db, "   ", fromID, toID, 100,
	)
	if err == nil {
		t.Fatal("expected empty idempotency key to fail")
	}

	if got := getBalance(t, db, fromID); got != 1000 {
		t.Errorf("sender balance changed unexpectedly: %d", got)
	}
}

