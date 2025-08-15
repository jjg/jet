package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	yesterday := time.Now().Add(-time.Hour * 24)
	yesterDate := yesterday.Format("2006-01-02")
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dataDir := fmt.Sprintf("%s/🦘pouch-data/%s", home, yesterDate)

	fmt.Printf("\nActivity for %s\n", yesterDate)

	// Get a list of files in the data directory.
	files, err := os.ReadDir(dataDir)
	if err != nil {
		panic(err)
	}

	// Display each file.
	for _, file := range files {
		filename := file.Name()

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
