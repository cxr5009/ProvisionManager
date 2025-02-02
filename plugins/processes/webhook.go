// plugins/processes/webhook.go
package processes

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// WebhookProcess implements core.ProcessPlugin to execute a shell command.
type WebhookProcess struct {
	URL  string
	Body []byte
}

func (w *WebhookProcess) Name() string {
	return "Webhook Process"
}

func (w *WebhookProcess) Execute() error {
	resp, err := http.Post(w.URL, "application/json", bytes.NewBuffer(w.Body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Webhook Response: %s\n", string(body))
	return nil
}
