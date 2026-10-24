package database

import (
	"context"
	"database/sql"
	"os"
	"sync"
	"testing"
)

func setupAccountsTest(t *testing.T) (*sql.DB, func()) {
	t.Helper()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("Skipping integration test: DATABASE_URL is not set")
	}

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

	_, err = db.Exec(
		"DELETE FROM accounts WHERE name LIKE 'Day17 Test %'",
	)
	if err != nil {
		db.Close()
		t.Fatalf("failed to clean test accounts: %v", err)
	}

	cleanup := func() {
		_, err := db.Exec(
			"DELETE FROM accounts WHERE name LIKE 'Day17 Test %'",
		)
		if err != nil {
			t.Errorf("failed to clean test accounts: %v", err)
		}
		db.Close()
	}

	return db, cleanup
}

func TestTransferBalance_Commit_Integration(t *testing.T) {
	db, cleanup := setupAccountsTest(t)
	defer cleanup()

	var fromID, toID int

	err := db.QueryRow(`
		INSERT INTO accounts (name, balance)
		VALUES ('Day17 Test Sender', 1000)
		RETURNING id
	`).Scan(&fromID)
	if err != nil {
		t.Fatal(err)
	}

	err = db.QueryRow(`
		INSERT INTO accounts (name, balance)
		VALUES ('Day17 Test Receiver', 200)
		RETURNING id
	`).Scan(&toID)
	if err != nil {
		t.Fatal(err)
	}

	if err := TransferBalance(
		context.Background(), db, fromID, toID, 300,
	); err != nil {
		t.Fatalf("transfer failed: %v", err)
	}

	var senderBalance, receiverBalance int64

	err = db.QueryRow(
		"SELECT balance FROM accounts WHERE id = $1", fromID,
	).Scan(&senderBalance)
	if err != nil {
		t.Fatal(err)
	}

	err = db.QueryRow(
		"SELECT balance FROM accounts WHERE id = $1", toID,
	).Scan(&receiverBalance)
	if err != nil {
		t.Fatal(err)
	}

	if senderBalance != 700 {
		t.Errorf("expected sender balance 700, got %d", senderBalance)
	}
	if receiverBalance != 500 {
		t.Errorf("expected receiver balance 500, got %d", receiverBalance)
	}
}

func TestTransferBalance_Rollback_Integration(t *testing.T) {
	db, cleanup := setupAccountsTest(t)
	defer cleanup()

	var fromID int

	err := db.QueryRow(`
		INSERT INTO accounts (name, balance)
		VALUES ('Day17 Test Rollback Sender', 1000)
		RETURNING id
	`).Scan(&fromID)
	if err != nil {
		t.Fatal(err)
	}

	err = TransferBalance(
		context.Background(), db, fromID, -999, 300,
	)
	if err == nil {
		t.Fatal("expected transfer to fail for nonexistent destination")
	}

	var senderBalance int64
	err = db.QueryRow(
		"SELECT balance FROM accounts WHERE id = $1", fromID,
	).Scan(&senderBalance)
	if err != nil {
		t.Fatal(err)
	}

	if senderBalance != 1000 {
		t.Errorf(
			"expected sender balance 1000 after rollback, got %d",
			senderBalance,
		)
	}
}

func TestTransferBalance_Concurrent_Integration(t *testing.T) {
	db, cleanup := setupAccountsTest(t)
	defer cleanup()

	ctx := context.Background()

	var fromID, toID int

	err := db.QueryRow(`
		INSERT INTO accounts (name, balance)
		VALUES ('Day17 Test Concurrent Sender', 10000)
		RETURNING id
	`).Scan(&fromID)
	if err != nil {
		t.Fatal(err)
	}

	err = db.QueryRow(`
		INSERT INTO accounts (name, balance)
		VALUES ('Day17 Test Concurrent Receiver', 0)
		RETURNING id
	`).Scan(&toID)
	if err != nil {
		t.Fatal(err)
	}

	const (
		workers        = 10
		amount         = int64(100)
		initialBalance = int64(10000)
	)

	var wg sync.WaitGroup
	errs := make(chan error, workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			errs <- TransferBalance(ctx, db, fromID, toID, amount)
		}()
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			t.Errorf("concurrent transfer failed: %v", err)
		}
	}

	var senderBalance, receiverBalance int64

	err = db.QueryRow(
		"SELECT balance FROM accounts WHERE id = $1", fromID,
	).Scan(&senderBalance)
	if err != nil {
		t.Fatal(err)
	}

	err = db.QueryRow(
		"SELECT balance FROM accounts WHERE id = $1", toID,
	).Scan(&receiverBalance)
	if err != nil {
		t.Fatal(err)
	}

	expectedSender := initialBalance - workers*amount
	expectedReceiver := int64(workers) * amount

	if senderBalance != expectedSender {
		t.Errorf("expected sender balance %d, got %d",
			expectedSender, senderBalance)
	}

	if receiverBalance != expectedReceiver {
		t.Errorf("expected receiver balance %d, got %d",
			expectedReceiver, receiverBalance)
	}

	if senderBalance+receiverBalance != initialBalance {
		t.Errorf("money conservation failed: sender=%d receiver=%d",
			senderBalance, receiverBalance)
	}
}

func TestTransferBalance_OppositeDirections_Concurrent_Integration(t *testing.T) {
	db, cleanup := setupAccountsTest(t)
	defer cleanup()

	var accountA, accountB int

	err := db.QueryRow(`
		INSERT INTO accounts (name, balance)
		VALUES ('Day17 Test Opposite A', 10000)
		RETURNING id
	`).Scan(&accountA)
	if err != nil {
		t.Fatal(err)
	}

	err = db.QueryRow(`
		INSERT INTO accounts (name, balance)
		VALUES ('Day17 Test Opposite B', 10000)
		RETURNING id
	`).Scan(&accountB)
	if err != nil {
		t.Fatal(err)
	}

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
				context.Background(),
				db,
				fromID,
				toID,
				100,
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

	var balanceA, balanceB int64

	err = db.QueryRow(
		"SELECT balance FROM accounts WHERE id = $1", accountA,
	).Scan(&balanceA)
	if err != nil {
		t.Fatal(err)
	}

	err = db.QueryRow(
		"SELECT balance FROM accounts WHERE id = $1", accountB,
	).Scan(&balanceB)
	if err != nil {
		t.Fatal(err)
	}

	if balanceA != 10000 {
		t.Errorf("expected account A balance 10000, got %d", balanceA)
	}

	if balanceB != 10000 {
		t.Errorf("expected account B balance 10000, got %d", balanceB)
	}

	if balanceA+balanceB != 20000 {
		t.Errorf("money conservation failed: total balance = %d",
			balanceA+balanceB)
	}
}
