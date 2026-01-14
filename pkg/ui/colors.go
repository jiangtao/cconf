package ui

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	// Color presets
	Green  = color.New(color.FgGreen).SprintFunc()
	Yellow = color.New(color.FgYellow).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
	Blue   = color.New(color.FgBlue).SprintFunc()
	Cyan   = color.New(color.FgCyan).SprintFunc()
	White  = color.New(color.FgWhite).SprintFunc()
	Gray   = color.New(color.FgHiBlack).SprintFunc()

	// Bold colors
	BoldGreen  = color.New(color.FgGreen, color.Bold).SprintFunc()
	BoldYellow = color.New(color.FgYellow, color.Bold).SprintFunc()
	BoldRed    = color.New(color.FgRed, color.Bold).SprintFunc()
	BoldBlue   = color.New(color.FgBlue, color.Bold).SprintFunc()
)

// Success prints a success message
func Success(format string, args ...interface{}) {
	fmt.Printf("%s %s\n", Green("✓"), fmt.Sprintf(format, args...))
}

// Warning prints a warning message
func Warning(format string, args ...interface{}) {
	fmt.Printf("%s %s\n", Yellow("⚠"), fmt.Sprintf(format, args...))
}

// Error prints an error message
func Error(format string, args ...interface{}) {
	fmt.Printf("%s %s\n", Red("✗"), fmt.Sprintf(format, args...))
}

// Info prints an info message
func Info(format string, args ...interface{}) {
	fmt.Printf("%s %s\n", Blue("ℹ"), fmt.Sprintf(format, args...))
}

// Skipped prints a skipped message
func Skipped(format string, args ...interface{}) {
	fmt.Printf("%s %s\n", Gray("⊘"), fmt.Sprintf(format, args...))
}

// Title prints a title
func Title(format string, args ...interface{}) {
	fmt.Printf("\n%s %s\n\n", BoldBlue("==="), fmt.Sprintf(format, args...))
}

// Step prints a step message
func Step(step, total int, format string, args ...interface{}) {
	fmt.Printf("%s %s\n", Cyan(fmt.Sprintf("[%d/%d]", step, total)), fmt.Sprintf(format, args...))
}

// Println prints colored text
func Println(colorFn func(...interface{}) string, format string, args ...interface{}) {
	fmt.Println(colorFn(fmt.Sprintf(format, args...)))
}

// Printf prints formatted colored text
func Printf(colorFn func(...interface{}) string, format string, args ...interface{}) {
	fmt.Print(colorFn(fmt.Sprintf(format, args...)))
}
