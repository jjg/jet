package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"time"
)

func getDataDir(t time.Time) string {

	// Make sure we have a working data dir.
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dataDir := fmt.Sprintf("%s/pouch-data/%s", home, t.Format("2006-01-02"))

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

func storeList(list []listItem) {

	// The only list we currently store is today's list
	dataDir := getDataDir(time.Now())

	// Create or update the file for the specified journal entry.
	filename := fmt.Sprintf("%s/hop.txt", dataDir)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Write it to disk.
	w := bufio.NewWriter(f)
	for _, item := range list {
		line := fmt.Sprintf("%s %s", item.status, item.description)
		if _, err := w.WriteString(line); err != nil {
			panic(err)
		}
		if _, err := w.WriteString("\n"); err != nil {
			panic(err)
		}
	}
	w.Flush()
}

func getList(sweetDir string, dayDir string) {
	filename := fmt.Sprintf("%s/%s/hop.txt", sweetDir, dayDir)
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

		// TODO: Load list instead of just printing it.
		fmt.Printf("%s", buf[:n])
	}
}

type listItem struct {
	status      string
	description string
}

func displayList(list []listItem) {
	for i, item := range list {
		fmt.Printf("%d. %s %s\n", i+1, item.status, item.description)
	}
}

func main() {

	fmt.Println("this is hop")

	//t := time.Now()
	//todayFile := getListFile(t)

	//sweetDir := getSweetDir()
	//dayDir := t.Format("2006-01-02")

	//getList(sweetDir, dayDir)

	// TODO: Import any incomplete items from yesterday's list.
	// TODO: Try to open today's file and if it doesn't exist, create it.

	// TODO: Load the list
	// TODO: Should this be a map?
	var list []listItem

	//item := listItem{status: "*", description: "Make the hop tool"}
	//list = append(list, item)

	// Display the list
	displayList(list)

	// Prompt for command (add/remove/complete/quit)
	var cmd string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("(a)dd, (r)emove, (c)omplete, (q)uit)> ")
		scanner.Scan()
		cmd = scanner.Text()
		if cmd == "a" || cmd == "r" || cmd == "c" {
			switch cmd {
			case "a":

				// Add an item to the list
				fmt.Print("> ")
				item := listItem{}
				item.status = "*"
				scanner.Scan()
				item.description = scanner.Text()
				list = append(list, item)

				// Save the list
				storeList(list)

				displayList(list)
			case "r":

				// Remove an item from the list
				fmt.Print("Item number to remove > ")
				scanner.Scan()
				itemNumber, err := strconv.Atoi(scanner.Text())
				if err != nil {
					panic(err)
				}
				// Adjust number for offset
				itemNumber--
				list = append(list[:itemNumber], list[itemNumber+1:]...)
				displayList(list)
			case "c":

				// Mark an item completed
				// TODO: This item selection logic should go in its own function.
				fmt.Print("Item number to complete > ")
				scanner.Scan()
				itemNumber, err := strconv.Atoi(scanner.Text())
				if err != nil {
					panic(err)
				}
				// Adjust number for offset
				itemNumber--
				list[itemNumber].status = "X"
				displayList(list)
			}
		}
		if cmd == "q" {
			break
		}
	}
}
