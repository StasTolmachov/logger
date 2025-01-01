
# Logrus Custom Logger

This package provides a custom logger based on the popular [Logrus](https://github.com/sirupsen/logrus) library. The logger includes colorized output and file/function caller information for enhanced debugging.

## Features

- **Colorized Output**: Different log levels are displayed in distinct colors for better readability.
- **Caller Information**: Logs include the file name, line number, and function name where the log was generated.
- **Customizable Formatter**: Built-in formatter (`MyFormatter`) for terminal-friendly log output.

## Installation

Use `go get` to install the package:

```bash
go get github.com/StasTolmachov/logger/logrus
```

## Usage

Here’s an example of how to use the custom logger:

```go
package main

import (
	"errors"
	"github.com/StasTolmachov/logger/logrus"
)

func main() {
	// Initialize the logger
	logrus.MakeLogger()

	// Log messages
	logrus.Log.Info("This is an info message")
	logrus.Log.Warn("This is a warning message")
	logrus.Log.Error("This is an error message")

	// Log an error with details
	err := errors.New("something went wrong")
	logrus.Log.Errorf("Error occurred: %v", err)
}
```

### Example Output (Terminal)

```plaintext
[INFO]  | main.main       | main.go:10 | This is an info message
[WARN]  | main.main       | main.go:11 | This is a warning message
[ERROR] | main.main       | main.go:12 | This is an error message
[ERROR] | main.main       | main.go:15 | Error occurred: something went wrong
```

## API Documentation

### Functions

#### `MakeLogger()`

Initializes the global logger (`logrus.Log`) with the following settings:
- Output: Terminal (`os.Stdout`).
- Formatter: Custom `MyFormatter` for colorized logs.
- Caller Information: Enabled (`SetReportCaller(true)`).

### Types

#### `MyFormatter`

A custom log formatter that adds the following features:
- **Colorized log levels**:
    - `INFO`: Green
    - `WARN`: Yellow
    - `ERROR`: Red
    - Other levels have distinct colors as well.
- **Caller Information**: Displays the file name, line number, and function name in each log entry.

## Configuration

You can extend the logger by modifying its global instance `logrus.Log` after initialization. For example, to change the log level:

```go
logrus.Log.SetLevel(logrus.DebugLevel)
```

To add additional outputs (e.g., a file):

```go
file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err == nil {
    logrus.Log.SetOutput(io.MultiWriter(os.Stdout, file))
}
```

## Dependencies

This package relies on the following libraries:
- [Logrus](https://github.com/sirupsen/logrus): A structured logger for Go.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
