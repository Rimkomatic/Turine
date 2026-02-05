package ui

import (
	"fmt"
	"strconv"
	"strings"
)

func hexToRGB(hex string) string {
	hex = strings.TrimPrefix(hex, "#")

	r, _ := strconv.ParseInt(hex[0:2], 16, 0)
	g, _ := strconv.ParseInt(hex[2:4], 16, 0)
	b, _ := strconv.ParseInt(hex[4:6], 16, 0)

	return fmt.Sprintf("%d;%d;%d", r, g, b)
}

func Color(hex string, text string) string {
	return fmt.Sprintf("\x1b[38;2;%sm%s\x1b[0m", hexToRGB(hex), text)
}

func Error(msg string) {
	fmt.Println(Color("#FF4C4C", "✗ "+msg))
}
