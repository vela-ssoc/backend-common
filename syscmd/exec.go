package syscmd

import (
	"context"
	"os/exec"
	"time"
)

func ExecTimeout(timeout time.Duration, name string, args ...string) *Result {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return Exec(ctx, name, args...)
}

func Exec(ctx context.Context, name string, args ...string) *Result {
	stdout := limited(100 * 1024)
	stderr := limited(100 * 1024)
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout, cmd.Stderr = stdout, stderr

	start := time.Now()
	err := cmd.Run()
	du := time.Since(start)

	ret := &Result{
		Duration: du,
		Stdout:   stdout.Bytes(),
		Stderr:   stderr.Bytes(),
	}
	if err != nil {
		ret.Error = err.Error()
	}
	if state := cmd.ProcessState; state != nil {
		ret.PID = state.Pid()
	}

	return ret
}

type Result struct {
	PID      int           `json:"pid,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
	Error    string        `json:"error,omitempty"`
	Stdout   []byte        `json:"stdout,omitempty"`
	Stderr   []byte        `json:"stderr,omitempty"`
}
