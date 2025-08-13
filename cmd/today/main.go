package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
)

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

func showEntry(journalDir string, entryName string) {
	filename := fmt.Sprintf("%s/%s.jet.txt", journalDir, entryName)
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

func main() {
	t := time.Now()
	journalDir := getJournalDir()

	fmt.Println("its today")
	entryName := t.Format("2006-01-02")
	showEntry(journalDir, entryName)
}
