package pool

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestNewCommandPool(t *testing.T) {
	p := NewCommandPool(context.Background(), 30*time.Second)
	if p == nil {
		t.Fatal("expected non-nil pool")
	}
}

func TestExecuteCommand_Echo(t *testing.T) {
	p := NewCommandPool(context.Background(), 30*time.Second)
	out, err := p.ExecuteCommand("echo", "hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(string(out), "hello") {
		t.Fatalf("expected output to contain 'hello', got %q", string(out))
	}
}

func TestExecuteCommand_InvalidCommand(t *testing.T) {
	p := NewCommandPool(context.Background(), 30*time.Second)
	_, err := p.ExecuteCommand("nonexistent_binary_xyz_12345")
	if err == nil {
		t.Fatal("expected error for nonexistent command")
	}
}

func TestExecuteCommand_ContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	p := NewCommandPool(ctx, 30*time.Second)
	start := time.Now()
	_, err := p.ExecuteCommand("sleep", "5")
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected timeout error")
	}
	if elapsed > 2*time.Second {
		t.Fatalf("expected fast timeout, took %v", elapsed)
	}
}

func TestExecuteCommand_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	p := NewCommandPool(ctx, 30*time.Second)
	_, err := p.ExecuteCommand("echo", "test")
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
}
