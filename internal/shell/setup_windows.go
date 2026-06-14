//go:build windows

package shell

import (
	"os"
	"os/exec"
)

func setup(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}

func terminate() {
	// do nothing
}
