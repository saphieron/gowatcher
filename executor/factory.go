package executor

import (
	"context"
	"os/exec"
)

// Just wraps the os/exec Command factory to clarify the link of CommandExecutor to the existing code
func NewCommandExecutor(name string, arg ...string) CommandExecutor {
	return exec.Command(name, arg...)
}

// Just wraps the os/exec Command factory to clarify the link of CommandExecutor to the existing code, but with a context object
func NewCommandExecutorWithContext(ctx context.Context, name string, arg ...string) CommandExecutor {
	return exec.CommandContext(ctx, name, arg...)
}
