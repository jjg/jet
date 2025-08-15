package main

import (
	"bufio"
	"errors"
	"fmt"
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

func storeList(list []listItem, t time.Time) {
	dataDir := getDataDir(t)
	// Create or overwrite the file.
	filename := fmt.Sprintf("%s/hop.txt", dataDir)
	f, err := os.Create(filename)
	// TODO: Set permissions on file to owner only.
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
	var list []listItem
	dataDir := getDataDir(t)
	filename := fmt.Sprintf("%s/hop.txt", dataDir)
	f, err := os.Open(filename)
	if err != nil {
		// It's OK to return an empty list if no file exists.
		return list
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for {
		more := scanner.Scan()
		itemText := scanner.Text()
		if !more {
			break
		}
		// TODO: There must be a better way to do this...
		itemWords := strings.Fields(itemText)
		if len(itemWords) > 1 {
			item := listItem{status: itemWords[0], description: strings.Join(itemWords[1:], " ")}
			list = append(list, item)
		}
	}
	return list
}

func getCompleted(list []listItem) []listItem {
	var completed []listItem
	for _, item := range list {
		if item.status == "X" {
			completed = append(completed, item)
		}
	}
	return completed
}

func getIncomplete(list []listItem) []listItem {
	var incomplete []listItem
	for _, item := range list {
		if item.status == "*" {
			incomplete = append(incomplete, item)
		}
	}
	return incomplete
}

func displayList(list []listItem) {
	for i, item := range list {
		fmt.Printf("%d. %s %s\n", i+1, item.status, item.description)
	}
}

func main() {

	today := time.Now()

	// Import any incomplete items from yesterday's list.
	yesterday := today.Add(-time.Hour * 24)
	completed := getCompleted(loadList(yesterday))
	incomplete := getIncomplete(loadList(yesterday))
	storeList(completed, yesterday)

	// Load today's list
	list := append(incomplete, loadList(today)...)
	storeList(list, today)

	// Display it.
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
				// TODO: Consider making this the default action (just start typing).
				// TODO: Don't add duplicate items to the list.
				fmt.Print("> ")
				item := listItem{}
				item.status = "*"
				scanner.Scan()
				item.description = scanner.Text()
				list = append(list, item)
				storeList(list, today)
				displayList(list)
			case "r":
				// Remove an item from the list
				fmt.Print("Item number to remove > ")
				scanner.Scan()
				itemNumber, err := strconv.Atoi(scanner.Text())
				if err != nil {
					panic(err)
				}
				itemNumber--
				list = append(list[:itemNumber], list[itemNumber+1:]...)
				storeList(list, today)
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
				itemNumber--
				list[itemNumber].status = "X"
				storeList(list, today)
				displayList(list)
			}
		}
		if cmd == "q" {
			break
		}
	}
}
