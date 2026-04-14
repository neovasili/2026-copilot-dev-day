package tools

import (
	"fmt"
)

const (
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	magenta = "\033[35m"
	reset   = "\033[0m"
)

// Magenta formats a raw message in magenta color.
func Magenta(message string) string {
	return fmt.Sprintf(magenta+"%s"+reset, message)
}

// Red formats a raw message in red color.
func Red(message string) string {
	return fmt.Sprintf(red+"%s"+reset, message)
}

// Green formats a raw message in green color.
func Green(message string) string {
	return fmt.Sprintf(green+"%s"+reset, message)
}

// Yellow formats a raw message in yellow color.
func Yellow(message string) string {
	return fmt.Sprintf(yellow+"%s"+reset, message)
}
