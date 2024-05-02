package colorful_logging

import (
	"fmt"
	"strings"
)

var (
	colors = map[string]string{
		"RESET":       "\033[0m",
		"UNDERLINE":   "\033[4m",
		"BLACK":       "\033[30m",
		"DARK RED":    "\033[31m",
		"DARK GREEN":  "\033[32m",
		"DARK YELLOW": "\x1b[33m",
		"DARK BLUE":   "\033[34m",
		"DARK PINK":   "\033[35m",
		"DARK CYAN":   "\033[36m",

		"GRAY": "\033[37m",

		"BRIGHT BLACK": "\033[90m",
		"RED":          "\033[91m",
		"GREEN":        "\033[92m",
		"YELLOW":       "\033[93m",
		"BLUE":         "\033[94m",
		"PINK":         "\033[95m",
		"CYAN":         "\033[96m",
		"BOLD":         "\033[97m",
	}
)

// ColorPrintln Is ColorPrint but generates new line
func ColorPrintln(a ...any) {
	ColorPrint(a)
	fmt.Print("\n")
}

// ColorPrint simply print anything with given color. last argument should be color name, otherwise the color name will print out.
// usage guide: ColorPrint(someObject, anotherObject, <color name>)
// for example: ColorPrint(someObject, anotherObject, "cyan")
// will print out then just like fmt.Print, but in given color.
func ColorPrint(a ...any) {
	var color string
	if c, ok := a[len(a)-1].(string); ok {
		color = GetColor(c)
		if color != "" {
			a = a[:len(a)-1]
		}
	}

	finalText := fmt.Sprint(a)
	finalText = finalText[1:]
	finalText = finalText[:len(finalText)-1]
	finalText = color + finalText + GetColor("reset")
	fmt.Println(finalText)
}

// GetColor Possible values for is as follows: ['red', 'green', 'yellow', 'blue',
// 'pink', 'cyan', 'gray', 'black', 'dark red', 'dark green', 'dark yellow',
// 'dark blue', 'dark pink', 'dark cyan', 'bright black']
func GetColor(color string) string {
	color = strings.ToUpper(color)
	c := colors[color]
	if len(c) == 0 {
		panic("Color not found")
	}
	return c
}
