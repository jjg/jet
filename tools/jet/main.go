package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/term"
	"io/fs"
	"os"
	"time"
)

func getDataDir(t time.Time) string {
	// Make sure we have a working data dir.
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dataDir := fmt.Sprintf("%s/🦘pouch-data/%s", home, t.Format("2006-01-02"))
	// Check if data dir exists and if not, create it.
	_, err = os.Stat(dataDir)
	if errors.Is(err, fs.ErrNotExist) {
		// Try to create the datadir.
		if err := os.MkdirAll(dataDir, 0700); err != nil {
			panic(err)
		}
	}
	return dataDir
}

func storeEntry(entry []string, t time.Time) {
	dataDir := getDataDir(t)
	// Create or update the file for the specified journal entry.
	filename := fmt.Sprintf("%s/📓.jet.txt", dataDir)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
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

func drawHeader() {
	// Get the terminal dimensions (ignore height since we don't use it).
	termWidth, _, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	// Draw ruler.
	rulerWidth := termWidth - 4
	fmt.Print("  |")
	for i := 1; i < rulerWidth; i++ {
		if i%10 == 0 {
			fmt.Printf("%d", i)
			i = i + len(fmt.Sprintf("%d", i)) - 1
		} else {
			fmt.Print("-")
		}
	}
	fmt.Print("|\n")
}

func getInput(interactive bool) []string {
	if interactive {
		drawHeader()
	}
	entry := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if interactive {
			fmt.Print("> ")
		}
		more := scanner.Scan()
		text := scanner.Text()
		entry = append(entry, text)
		// Break on double linefeed if in interactive mode.
		if interactive {
			if len(text) == 0 {
				break
			}
		} else {
			if !more {
				break
			}
		}
	}
	return entry
}

func main() {
	today := time.Now()
	// Look for subcommands on the command line.
	subCommand := "newInteractiveEntry"
	args := os.Args
	if len(args) > 1 {
		subCommand = args[1]
	}
	switch subCommand {
	case "help":
		fmt.Println("A reliable focused writing tool.")
		fmt.Println("")
		fmt.Println("Journal entries for each day will be stored in a subdirectory")
		fmt.Println("in the 'pouch-data' in your home directory.")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("jet - Create a new journal entry in interactive mode")
		fmt.Println("cat foo.txt | jet - Create a new journal entry using the output from cat")
		fmt.Println("")
		fmt.Println("See also:")
		fmt.Println("today, yesterday")
		fmt.Println("")
	default:
		// If no subcommand is provided, create a new entry.
		entry := getInput(term.IsTerminal(0))
		storeEntry(entry, today)
	}
}
