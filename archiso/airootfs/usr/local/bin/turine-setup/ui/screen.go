package ui

import "fmt"

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func ClearLine() {
	fmt.Print("\033[2K\r")
}

func HideCursor() {
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?25h")
}
