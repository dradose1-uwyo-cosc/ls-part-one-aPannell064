package functions

import (
	"os"
	"strings"
)

// removes any hidden files from the dir listing, note not exported
func dirFilter(entries []os.DirEntry) []os.DirEntry {
	n := 0
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), ".") {
			entries[n] = entry
			n++
		}
	}
	return entries
}
