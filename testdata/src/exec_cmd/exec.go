package exec

import (
	"context"
	"os/exec"
)

func _() {
	ctx := context.Background()

	exec.Command("ls", "-l") // want `os/exec.Command must not be called. use os/exec.CommandContext`

	exec.CommandContext(ctx, "ls", "-l")
}
