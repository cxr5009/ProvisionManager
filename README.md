# Provision Manager TUI

Provision Manager TUI is a modular, pluggable terminal user interface (TUI) application written in Go using [Bubble Tea](https://github.com/charmbracelet/bubbletea). It is designed to help you manage system requirements, select programs, and execute a series of processes—all via a clean, keyboard-driven interface.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Extending the Application](#extending-the-application)
  - [Adding a Requirement Plugin](#adding-a-requirement-plugin)
  - [Adding a Process Plugin](#adding-a-process-plugin)
  - [Adding a Program Plugin](#adding-a-program-plugin)
- [Contributing](#contributing)
- [License](#license)

## Overview

This application is divided into three major steps:

1. **Requirements Checking:**  
   Uses pluggable requirement plugins (e.g., ADB, LSUSB) to verify that system prerequisites are met. Each plugin implements methods to check, install, or update a requirement.

2. **Program Selection:**  
   Lets users choose a program from a list. Each program is a plugin that specifies a series of processes (steps) to run.

3. **Program Execution:**  
   Executes a series of processes for the selected program. Process plugins perform actions such as running shell commands or making HTTP requests.

All three portions of the framework are hookable and pluggable, meaning you can easily extend or modify the application to suit your needs.

## Features

- **Modular Design:**  
  Easily add new requirement, process, or program plugins without changing core logic.
  
- **TUI with Bubble Tea:**  
  Enjoy an intuitive, keyboard-driven interface.
  
- **Pluggable Architecture:**  
  Extend functionality by simply creating new plugins that conform to defined interfaces.

## Project Structure
.
 ├── README.md
 ├── go.mod
 ├── go.sum
 ├── main.go
 ├── core
 │ └── interfaces.go # Shared interfaces for plugins
 ├── plugins
 │ ├── reqs
 │ │ ├── adb.go # Example: ADB requirement plugin
 │ │ └── lsusb.go # Example: LSUSB requirement plugin
 │ ├── programs
 │ │ └── sample_program.go # Example: A program plugin
 │ └── processes
 │ ├── shell_command.go # Example: Process plugin for shell commands
 │ └── webhook.go # Example: Process plugin for HTTP requests
 └── tui
   └── model.go # Bubble Tea TUI model
> **Note:** Shared interfaces are defined in the `core` package to avoid circular dependencies and prevent importing from the `main` package.
## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/cxr5009/provisionmanager.git
    cd provisionmanager
2. **Download dependencies:**

    ```bash
    go mod tidy
3. **Build the application:**

    ```bash
    go build -o provisionmanager
4. **Run the application:**

    ```bash
    ./provisionmanager
## Usage
When you run the application, you will be guided through three steps via the TUI:

1. **Requirements:**
The application will check if required software (e.g., ADB) is installed. Press `[1]` to initiate requirement checks.

2. **Program Selection:**
Once requirements are verified, you can choose a program to run. Press `[2]` to select a program.

3. **Execution:**
After selecting a program, press `[3]` to execute its processes (for example, executing a shell command or sending a webhook).

Press `[q]` at any time to quit the application.

## **Extending the Application**
The architecture is designed to be extended. Below are guidelines for adding new plugins.

### **Adding a Requirement Plugin**
1. **Create a new file in `plugins/reqs/`**, for example `docker.go`.
2. **Implement the `RequirementPlugin` interface defined** in `core/interfaces.go`:

    ```go
    // plugins/reqs/docker.go
    package reqs

    import (
        "fmt"
        "os/exec"
        "errors"

        "github.com/cxr5009/provisionmanager/core"
    )

    // DockerPlugin implements core.RequirementPlugin.
    type DockerPlugin struct{}

    func (p *DockerPlugin) Name() string {
        return "Docker"
    }

    func (p *DockerPlugin) Check() error {
        // Example: check Docker version
        _, err := exec.Command("docker", "version").Output()
        if err != nil {
            return errors.New("docker not found")
        }
        return nil
    }

    func (p *DockerPlugin) Install() error {
        fmt.Println("Installing Docker...")
        // Add your installation logic here.
        return nil
    }

    func (p *DockerPlugin) Update() error {
        fmt.Println("Updating Docker...")
        // Add your update logic here.
        return nil
    }
3. **Register the plugin:**
You can create a plugin registry in your main.go (or a separate registry file) to load all available requirement plugins.

### **Adding a Process Plugin**
1. **Create a new file in `plugins/processes/`**, for example `webhook.go`
2. **Implment the `ProcessPlugin` interface** defined in `core/interfaces.go`:

    ```go
    // plugins/processes/webhook.go
    package processes

    import (
        "fmt"
        "net/http"
        "bytes"
        "io/ioutil"

        "github.com/cxr5009/provisionmanager/core"
    )

    // WebhookProcess implements core.ProcessPlugin.
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

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return err
        }
        fmt.Printf("Webhook Response: %s\n", string(body))
        return nil
    }

### Adding a Program Plugin
1. **Create a new file in `plugins/programs/`**, for example `custom_program.go`.
2. **Implement the `ProgramPlugin` interface** defined in `core/interfaces.go`.

    ```go
    // plugins/programs/custom_program.go
    package programs

    import (
        "github.com/cxr5009/provisionmanager/core"
        "github.com/cxr5009/provisionmanager/plugins/processes"
    )

    // CustomProgram implements core.ProgramPlugin.
    type CustomProgram struct{}

    func (cp *CustomProgram) Name() string {
        return "Custom Program"
    }

    func (cp *CustomProgram) Processes() []core.ProcessPlugin {
        return []core.ProcessPlugin{
            // Example: A shell command process
            &processes.ShellCommandProcess{
                Command: "echo",
                Args:    []string{"Hello from Custom Program!"},
            },
            // You can add more processes here.
        }
    }
3. **Update the program selection register** (if applicable) so  the TUI displayes the new program as an option.

## Contributing
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a feature branch.
3. Write tests and update documentation.
4. Submit a pull request describing your changes.

For major changes, please open an issue first to discuss what you would like to change.

## License
Distributed under the MIT License. See LICENSE for more information.