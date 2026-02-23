package functions

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
)

// Performs ls on a directory
func lsDir(w io.Writer, useColor bool, dir string) {
	// Get files in directory
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Filter entries
	entries = dirFilter(entries)

	for _, entry := range entries {
		// Get name of each entry
		name := entry.Name()
		if useColor {
			// Get file info
			info, err := os.Lstat(filepath.Join(dir, name))
			if err != nil {
				fmt.Fprint(os.Stderr, err)
			}
			mode := info.Mode()

			if info.IsDir() {
				// Print directories in blue
				Blue.ColorPrint(w, name)
			} else if mode.IsRegular() && mode&0111 != 0 {
				// Print files/executables in green
				Green.ColorPrint(w, name)
			} else {
				// Print anything else in default color
				Reset.ColorPrint(w, name)
			}
		} else {
			// Print in default color not using colors
			Reset.ColorPrint(w, name)
		}
	}
}

// A simple ls for when no flags provided
func SimpleLS(w io.Writer, args []string, useColor bool) {
	if len(args) == 0 {
		// Perform ls on current directory if no arguments are given
		lsDir(w, useColor, ".")
	} else {
		// Split args into files and directories
		var files []string
		var dirs []string
		for _, arg := range args {
			info, err := os.Lstat(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
			}
			if info.IsDir() {
				dirs = append(dirs, arg)
			} else {
				files = append(files, arg)
			}
		}

		// Explicitely sort both sets of data (doesn't usually do anything)
		sort.Strings(dirs)
		sort.Strings(files)

		// Perform ls on files
		for _, file := range files {
			if useColor {
				// Get file info
				info, err := os.Lstat(file)
				if err != nil {
					fmt.Fprint(os.Stderr, err)
				}
				mode := info.Mode()

				if mode.IsRegular() && mode&0111 != 0 {
					// Print files/executables in green
					Green.ColorPrint(w, file)
				} else {
					// Print anything else in default color
					Reset.ColorPrint(w, file)
				}
			} else {
				// Print in default color not using colors
				Reset.ColorPrint(w, file)
			}

		}

		// Perform ls on each directory
		for _, dir := range dirs {
			if len(args) > 1 {
				// Add a header in default color if there were multiple arguments
				Reset.ColorPrint(w, "\n"+dir+": ")
			}
			lsDir(w, useColor, dir)
		}
	}
}
