package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	today := time.Now()
	toDate := today.Format("2006-01-02")
	home, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	dataDir := fmt.Sprintf("%s/🦘pouch-data/%s", home, toDate)

	fmt.Printf("\nActivity for %s\n", toDate)

	// Get a list of files in the data directory.
	files, err := os.ReadDir(dataDir)

	if err != nil {
		panic(err)
	}

	// Display each file.

	for _, file := range files {
		filename := file.Name()

		// TODO: Find a way to trim the extension w/o breaking the emoji.
		fmt.Printf("\n--- %s ---\n", filename)

		f, err := os.Open(fmt.Sprintf("%s/%s", dataDir, filename))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for {
			more := scanner.Scan()
			if !more {
				break
			}
			fmt.Println(scanner.Text())
		}
	}
}
