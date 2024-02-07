package logger

import (
	"fmt"
)

const (
	ColorRed    = "\033[91m"
	ColorGreen  = "\033[92m"
	ColorYellow = "\033[93m"
	ColorBlue   = "\033[94m"
	ColorReset  = "\033[0m"
)

func Log(a ...interface{}) {
	fmt.Println(a...)
}

func Print(a ...interface{}) {
	fmt.Print(a...)
}

// LogError logs an error message
func LogError(err error) {
	fmt.Println(ColorRed+"ERROR:", err, ColorReset)
}

// LogSuccess logs a success message
func LogSuccess(message string) {
	fmt.Println(ColorGreen + message + ColorReset)
}

// LogWarning logs a warning message
func LogWarning(message string) {
	fmt.Println(ColorYellow + message + ColorReset)
}
