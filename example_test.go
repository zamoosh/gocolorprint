// Developed by ZAMOOSH

package colorful_logging

import (
	"fmt"
	"time"
)

func ExampleColorPrint() {
	text := "hello world!"
	now := time.Now().UTC()

	ColorPrint(text, now, "cyan")
}

func ExampleColorPrintln() {
	text := "hello world!"
	now := time.Now().UTC()

	ColorPrintln(text, now, "cyan")
}

// DO NOT FORGET TO END YOUR PRINT WITH "reset" to prevent continuing coloring!
func ExampleGetColor() {
	red := GetColor("red")
	reset := GetColor("reset")
	text := "hello world!"

	fmt.Println(red, text, reset)
}
