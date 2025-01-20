# Custom Logger in Go

A simple yet flexible logger written in Go, designed to provide custom logging with optional emojis, timestamps, and color-coded messages for better readability. The logger includes predefined methods for different message types (`Success`, `Error`, `Warning`, `Info`) and allows creating custom message types.

## Features

- **Predefined message types**: Success, Error, Warning, and Info.
- **Customizable logging**: Create your own log types with specific emojis and colors.
- **Optional emojis and timestamps**: Toggle these features as needed.
- **Color-coded output**: Improve log readability in the terminal.

## Usage

Here's a quick example of how to use this logger:

```go
package main

import (
	"github.com/evertonmsantos/signale"
	"github.com/fatih/color"
)

func main() {
	logger := signale.New(true, true) // Emojis and timestamps enabled

	logger.Success("Operation completed successfully!")
	logger.Error("An error occurred during the process.")
	logger.Warning("Disk space is running low.")
	logger.Info("System is starting up.")

	// Custom log type
	logger.Custom("debug", "🐞", color.FgMagenta, "This is a debug log.")
}
```

## Installation

1. Install the package:
   ```bash
   go get github.com/evertonmsantos/signale
   ```

2. Run the example by creating a file (e.g., `main.go`) and using the code snippet provided above:
   ```bash
   go run main.go
   ```

## Contributions

Feel free to open issues or submit pull requests if you have suggestions or improvements!
