package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"
)

func main() {

	// TODO: Provide better erors (don't just panic() all the time).
	// TODO: Allow journal dir to be cusomized.
	// TODO: Provide some help/instructions.

	// Make sure we have a working journal dir before they write anything.
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	journalDir := fmt.Sprintf("%s/jet-journal", home)

	// Check if journal dir exists and if not, create it.
	_, err = os.Stat(journalDir)
	if errors.Is(err, fs.ErrNotExist) {

		// Try to create the journal dir.
		// TODO: Maybe notify (and confirm?) before doing this?
		if err := os.MkdirAll(journalDir, 0750); err != nil {
			panic(err)
		}
	}

	// Get the name for today's entry
	t := time.Now()
	dateString := t.Format("2006-01-02")

	// Draw ruler
	fmt.Printf("                                                                       %s\n", dateString)
	fmt.Println("  |--------|---------|---------|---------|---------|---------|---------|---------|")

	// Read input
	entry := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		// Read until EOF (blank line)
		if len(text) != 0 {
			entry = append(entry, text)
		} else {
			break
		}
	}

	// Create or update today's journal file
	filename := fmt.Sprintf("%s/%s.txt", journalDir, dateString)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Write it to disk.
	w := bufio.NewWriter(f)
	for _, line := range entry {
		if _, err := w.WriteString(line); err != nil {
			panic(err)
		}
		if _, err := w.WriteString("\n"); err != nil {
			panic(err)
		}
	}
	w.Flush()
}
