package database

import (
	"context"
	"errors"
	"strings"
	"testing"

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

	if !strings.Contains(err.Error(), "canceled") {
		t.Errorf("expected context cancellation error, got: %v", err)
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

	if attempts != maxTransferAttempts {
		t.Fatalf(
			"expected %d attempts, got %d",
			maxTransferAttempts,
			attempts,
		)
	}
}
