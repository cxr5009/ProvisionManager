package processes

import (
	"fmt"
)

// WebhookProcess implements ProcessPlugin to execute a shell command.
type WebhookProcess struct {
	Command string
	Args    []string
}

func (s *WebhookProcess) Name() string {
	return "Webhook Command: " + s.Command
}

func (s *WebhookProcess) Execute() error {
	fmt.Printf("Executing: %s %v\n", s.Command, s.Args)
	// TODO: Implement webhook logic
	return nil
}
