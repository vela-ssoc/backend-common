package syscmd

import (
	"context"
	"os/exec"
	"strings"
	"time"
)

type Input struct {
	Cmd     string        `json:"cmd"     query:"cmd"     validate:"required,lte=255"`
	Args    []string      `json:"args"    query:"args"    validate:"lte=50"`
	Timeout time.Duration `json:"timeout" query:"timeout"`
}

type Output struct {
	Cmd      string   `json:"cmd"`
	Args     []string `json:"args"`
	PID      int      `json:"pid"`
	Duration string   `json:"duration"`
	Error    string   `json:"error"`
	Stdout   []byte   `json:"stdout"`
	Stderr   []byte   `json:"stderr"`
}

func Exec(input Input, strict bool) Output {
	return ExecWithContext(nil, input, strict)
}

func ExecWithContext(parent context.Context, input Input, strict bool) Output {
	if parent == nil {
		parent = context.Background()
	}
	timeout := input.Timeout
	if timeout < time.Second || timeout > 30*time.Second {
		timeout = 30 * time.Second
	}
	name, args := input.Cmd, input.Args
	// 非严格模式，增强用户易用性
	if !strict {
		asz := len(args)
		if asz == 0 {
			if split := strings.Split(name, " "); len(split) > 1 {
				name, args = split[0], split[1:]
			}
		} else if asz == 1 {
			args = strings.Split(args[0], " ")
		}
	}

	ctx, cancel := context.WithTimeout(parent, timeout)
	cmd := exec.CommandContext(ctx, name, args...)
	stdout := limited(100 * 1024)
	stderr := limited(100 * 1024)
	cmd.Stdout, cmd.Stderr = stdout, stderr

	start := time.Now()
	err := cmd.Run()
	du := time.Since(start)
	cancel()

	var pid int
	var cause string
	if err != nil {
		cause = err.Error()
	}
	if state := cmd.ProcessState; state != nil {
		pid = state.Pid()
	}

	return Output{
		Cmd:      name,
		Args:     args,
		PID:      pid,
		Duration: du.String(),
		Error:    cause,
		Stdout:   stdout.Bytes(),
		Stderr:   stderr.Bytes(),
	}
}
