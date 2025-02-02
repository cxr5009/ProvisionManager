package processes

import (
	"fmt"
	"os/exec"
)

// ShellCommandProcess implements ProcessPlugin to execute a shell command.
type ShellCommandProcess struct {
	Command string
	Args    []string
}

func (s *ShellCommandProcess) Name() string {
	return "Shell Command: " + s.Command
}

func (s *ShellCommandProcess) Execute() error {
	fmt.Printf("Executing: %s %v\n", s.Command, s.Args)
	cmd := exec.Command(s.Command, s.Args...)
	out, err := cmd.CombinedOutput()
	fmt.Printf("Output: %s\n", out)
	return err
}
