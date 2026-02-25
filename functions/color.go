package functions

import (
	"fmt"
	"io"
	"os"
)

// List of colors
type color string

const (
	Reset color = "\x1b[0m"
	Green color = "\x1b[32m"
	Blue  color = "\x1b[34m"
)

// Prints appropriate entries in color, or colorless if regular file
func (c color) ColorPrint(w io.Writer, s string) {
	_, err := w.Write([]byte(string(c) + s + string(Reset) + "\n"))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
