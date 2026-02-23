// Adam Pannell
// COSC 3750
// 02/26/2026

package main

import (
	"flag"
	"gols/functions"
	"os"
)

func main() {
	flag.Parse()
	functions.SimpleLS(os.Stdout, flag.Args(), functions.IsTerminal(os.Stdout))
}
