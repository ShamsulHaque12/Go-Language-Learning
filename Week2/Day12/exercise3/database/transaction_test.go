package database

import (
	"context"
	"strings"
	"testing"
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
