package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	// TODO: Look for settings (~/.config/jet/settings.json)
	journalDir := "/home/jason/journal"

	// Draw ruler
	fmt.Println("  |--------|---------|---------|---------|---------|---------|---------|---------|")

	// Read input
	entry := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		// Listen for EOF
		if len(text) != 0 {
			entry = append(entry, text)
		} else {
			break
		}
	}

	// Compute filename
	t := time.Now()
	filename := fmt.Sprintf("%s/%s.txt", journalDir, t.Format("2006-01-02"))

	// TODO: Check if journal directory exists and create if not

	// Create or update today's journal file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range entry {

		// TODO: See if there's a better way to add this linefeed
		formattedLine := fmt.Sprintf("%s\n", line)

		_, err := w.WriteString(formattedLine)
		if err != nil {
			panic(err)
		}
	}

	w.Flush()
}
