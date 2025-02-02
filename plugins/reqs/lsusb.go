package reqs

import (
	"errors"
	"fmt"
	"os/exec"
)

// LSUSBPlugin implements RequirementPlugin for lsusb.
type LSUSBPlugin struct{}

func (p *LSUSBPlugin) Name() string {
	return "ADB"
}

// Check if LSUSB is installed (for example, by checking its version).
func (p *LSUSBPlugin) Check() error {
	_, err := exec.Command("lsusb", "version").Output()
	if err != nil {
		return errors.New("LSUSB is not installed")
	}
	return nil
}

func (p *LSUSBPlugin) Install() error {
	fmt.Println("Installing LSUSB...")
	// TODO: installation process
	return nil
}

// Update LSUSB if needed.
func (p *LSUSBPlugin) Update() error {
	fmt.Println("Updating LSUSB...")
	// TODO: update process
	return nil
}
