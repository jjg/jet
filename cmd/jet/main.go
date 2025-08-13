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

func getJournalDir() string {

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
		if err := os.MkdirAll(journalDir, 0700); err != nil {
			panic(err)
		}
	}

	return journalDir
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

func storeEntry(journalDir string, entryName string, entry []string) {

	// Create or update the file for the specified journal entry.
	filename := fmt.Sprintf("%s/%s.jet.txt", journalDir, entryName)
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

func main() {
	t := time.Now()
	journalDir := getJournalDir()

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
		fmt.Println("Journal entries for each day will be stored in a directory")
		fmt.Println("named 'jet-journal' in your home directory.")
		fmt.Println("")
		fmt.Println("usage: jet (help, today, yesterday)")
		fmt.Println("jet\t\t- Create a new journal entry")
		fmt.Println("jet help\t- this message")
		fmt.Println("jet today\t- Show today's journal entries")
		fmt.Println("jet yesterday\t- Show yesterday's entries")
		fmt.Println("")
	default:

		// If no subcommand is provided, create a new entry.
		entryName := t.Format("2006-01-02")
		entry := getInput(term.IsTerminal(0))
		storeEntry(journalDir, entryName, entry)
	}
}
