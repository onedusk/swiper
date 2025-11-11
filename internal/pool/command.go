package pool

import (
	"context"
	"os/exec"
	"sync"
	"time"
)

// CommandPool manages a pool of reusable command executors
type CommandPool struct {
	mu       sync.Mutex
	cmdCache map[string]*exec.Cmd
	ctx      context.Context
}

// NewCommandPool creates a new command pool
func NewCommandPool(ctx context.Context) *CommandPool {
	return &CommandPool{
		cmdCache: make(map[string]*exec.Cmd),
		ctx:      ctx,
	}
}

// ExecuteCommand executes a command with timeout and caching
func (p *CommandPool) ExecuteCommand(name string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(p.ctx, 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.Output()
}