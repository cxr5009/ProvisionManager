package reqs

import (
	"errors"
	"fmt"
	"os/exec"
)

// ADBPlugin implements RequirementPlugin for ADB.
type ADBPlugin struct{}

func (p *ADBPlugin) Name() string {
	return "ADB"
}

// Check if ADB is installed (for example, by checking its version).
func (p *ADBPlugin) Check() error {
	_, err := exec.Command("adb", "version").Output()
	if err != nil {
		return errors.New("ADB is not installed")
	}
	return nil
}

func (p *ADBPlugin) Install() error {
	fmt.Println("Installing ADB...")
	// TODO: installation process
	return nil
}

// Update ADB if needed.
func (p *ADBPlugin) Update() error {
	fmt.Println("Updating ADB...")
	// TODO: update process
	return nil
}
