package log

import (
	"fmt"
	"os"
)

// Info prints an informational message.
func Info(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "[INFO] "+format+"\n", args...)
}

// Error prints an error message.
func Error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "[ERROR] "+format+"\n", args...)
}
