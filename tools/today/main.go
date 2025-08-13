package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func getJournalDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/jet-journal", home)
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
	entryName := t.Format("2006-01-02")
	showEntry(journalDir, entryName)
}
