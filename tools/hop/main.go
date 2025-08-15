package main

import (
	"bufio"
	"errors"
	"fmt"
	//"io"
	"io/fs"
	"os"
	"strconv"
	"strings"
	"time"
)

type listItem struct {
	status      string
	description string
}

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
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0660)
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

func loadList(t time.Time) []listItem {
	dataDir := getDataDir(t)
	filename := fmt.Sprintf("%s/hop.txt", dataDir)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var list []listItem
	scanner := bufio.NewScanner(f)
	for {
		more := scanner.Scan()
		itemText := scanner.Text()

		if !more {
			break
		}

		// TODO: There must be a better way to do this...
		itemWords := strings.Fields(itemText)
		item := listItem{status: itemWords[0], description: strings.Join(itemWords[1:], " ")}
		list = append(list, item)
	}

	return list
}

func displayList(list []listItem) {
	for i, item := range list {
		fmt.Printf("%d. %s %s\n", i+1, item.status, item.description)
	}
}

func main() {

	// TODO: Import any incomplete items from yesterday's list.

	// Load the list
	// TODO: Should this be a map?
	//var list []listItem
	t := time.Now()
	list := loadList(t)

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
