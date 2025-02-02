package reqs

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

// ADBPlugin implements core.RequirementPlugin for ADB.
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

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "choco install adb")
	case "darwin":
		cmd = exec.Command("brew", "install", "android-platform-tools")
	case "linux":
		cmd = exec.Command("sudo", "apt-get", "install", "-y", "adb")
	default:
		return errors.New("unsupported operating system")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to install ADB: %s\n%s", err, output)
	}

	fmt.Println("ADB installed successfully.")
	return nil
}

// Update ADB if needed.
func (p *ADBPlugin) Update() error {
	fmt.Println("Updating ADB...")

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "choco upgrade adb")
	case "darwin":
		cmd = exec.Command("brew", "upgrade", "android-platform-tools")
	case "linux":
		cmd = exec.Command("sudo", "apt-get", "update", "&&", "sudo", "apt-get", "upgrade", "-y", "adb")
	default:
		return errors.New("unsupported operating system")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to update ADB: %s", string(output))
	}

	fmt.Println("ADB updated successfully")
	return nil
}
