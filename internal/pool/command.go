package pool

import (
	"context"
	"os/exec"
	"time"
)

// CommandPool manages a pool of reusable command executors
type CommandPool struct {
	ctx     context.Context
	timeout time.Duration
}

// NewCommandPool creates a new command pool with the given timeout.
// If timeout <= 0, defaults to 30 seconds.
func NewCommandPool(ctx context.Context, timeout time.Duration) *CommandPool {
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	return &CommandPool{
		ctx:     ctx,
		timeout: timeout,
	}
}

// ExecuteCommand executes a command with the configured timeout
func (p *CommandPool) ExecuteCommand(name string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(p.ctx, p.timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.Output()
}