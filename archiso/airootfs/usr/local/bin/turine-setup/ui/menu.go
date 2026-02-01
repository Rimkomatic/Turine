package ui

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func SelectMenu(title string, items [][]string) int {
	fd := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState)
	defer ShowCursor()

	HideCursor()

	selected := 0
	buf := make([]byte, 3)

	for {
		ClearScreen()

		fmt.Println(title)
		fmt.Println()

		for i, item := range items {
			label := item[0]
			hex := item[1]

			r, g, b := RGBFromHex(hex)

			ClearLine()
			SetRGB(r, g, b)

			if i == selected {
				fmt.Printf("❯ %s", label)
			} else {
				fmt.Printf("  %s", label)
			}

			ResetColor()
			fmt.Println()
		}

		n, _ := os.Stdin.Read(buf)
		if n == 0 {
			continue
		}

		if buf[0] == 27 && n >= 3 && buf[1] == 91 {
			switch buf[2] {
			case 'A': // up
				if selected > 0 {
					selected--
				}
			case 'B': // down
				if selected < len(items)-1 {
					selected++
				}
			}
		} else if buf[0] == 10 || buf[0] == 13 { // Enter
			return selected
		}
	}
}

