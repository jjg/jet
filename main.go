package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/term"
	"io"
	"io/fs"
	"os"
	"time"
)

func drawHeader(t string) {

	// Get the terminal dimensions (ignore height since we don't use it).
	termWidth, _, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	// Draw header text right-justified.
	headerWidth := termWidth - len(t) - 1
	for i := 0; i < headerWidth; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("%s\n", t)

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
		if err := os.MkdirAll(journalDir, 0750); err != nil {
			panic(err)
		}
	}

	return journalDir
}

func getInput() []string {
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

	return entry
}

func showEntry(journalDir string, entryName string) {
	filename := fmt.Sprintf("%s/%s.txt", journalDir, entryName)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		fmt.Printf("%s", buf[:n])
	}
}

func storeEntry(journalDir string, entryName string, entry []string) {

	// Create or update today's journal file
	filename := fmt.Sprintf("%s/%s.txt", journalDir, entryName)
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

func main() {
	t := time.Now()

	// TODO: Provide better erors (don't just panic() all the time).
	// TODO: Allow journal dir to be cusomized?
	// TODO: Provide some help/instructions.
	// TODO: Allow text to be piped-in?

	journalDir := getJournalDir()

	// Look for subcommands on the command line
	subCommand := "newInteractiveEntry"
	args := os.Args
	if len(args) > 1 {
		subCommand = args[1]
	}

	// DEBUG
	//fmt.Println(subCommand)

	switch subCommand {
	case "today":

		// Show today's entries
		entryName := t.Format("2006-01-02")
		showEntry(journalDir, entryName)

	default:

		// If no subcommand is provided, create a new entry.
		entryName := t.Format("2006-01-02")

		// If we're in interactive mode, draw the header.
		if term.IsTerminal(0) {
			drawHeader(entryName)
		}
		entry := getInput()
		storeEntry(journalDir, entryName, entry)
	}
}
