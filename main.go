package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// TODO: Look for settings (~/.config/jet/settings.json)

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

	// DEBUG: Dump entry to console
	fmt.Println(entry)

	// TODO: Seems like `entry` might be missing the linefeeds...

	// TODO: Check for existing journal file for today

	// TODO: Create or update today's journal file
	f, err := os.Create("journal/today.txt")
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
