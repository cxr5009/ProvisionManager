package main

// --------------------
// Requirement Plugin Interface
// --------------------
type RequirementPlugin interface {
	// Name returns the name of the requirement (e.g., "ADB", "LSUSB")
	Name() string
	// Check verifies if the requirement is met.
	Check() error
	// Install attempts to install the requirement.
	Install() error
	// Update updates the requirement if needed.
	Update() error
}

// --------------------
// Process Plugin Interface
// --------------------
type ProcessPlugin interface {
	// Execute performs the process (e.g., shell command or webhook)
	Execute() error
	// Name returns the process name
	Name() string
}

// --------------------
// Program Plugin Interface
// --------------------
type ProgramPlugin interface {
	// Name returns the program name.
	Name() string
	// Processes returns the list of processes that compose the program.
	Processes() []ProcessPlugin
}
