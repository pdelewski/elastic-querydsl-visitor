package main

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func printColoredTextWithBackground(text, textColor, backgroundColor string) {
	// ANSI escape code for clearing the line and setting background color
	fmt.Printf("\033[K%s%s%s%s%s\n", backgroundColor, textColor, text, Reset, Reset)
}
