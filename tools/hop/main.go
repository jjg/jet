package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	//"time"
)

func getSweetDir() string {

	// Make sure we have a working journal dir before they write anything.
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	sweetDir := fmt.Sprintf("%s/sweet", home)

	// Check if sweet dir exists and if not, create it.
	_, err = os.Stat(sweetDir)
	if errors.Is(err, fs.ErrNotExist) {

		// Try to create the sweet dir.
		if err := os.MkdirAll(sweetDir, 0700); err != nil {
			panic(err)
		}
	}

	return sweetDir
}

func storeList(sweetDir string, dayDir string, list []string) {

	// Select or create today's directory.
	todayDir := fmt.Sprintf("%s/%s", sweetDir, dayDir, list)
	_, err := os.Stat(todayDir)
	if errors.Is(err, fs.ErrNotExist) {
		// Try to create the today dir.
		if err := os.MkdirAll(todayDir, 0700); err != nil {
			panic(err)
		}
	}

	// Create or update the file for the specified journal entry.
	filename := fmt.Sprintf("%s/hop.txt", todayDir)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Write it to disk.
	w := bufio.NewWriter(f)
	for _, line := range list {
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
		fmt.Printf("%d. %s\n", i+1, item.description)
	}
}

func main() {

	fmt.Println("this is hop")

	//t := time.Now()
	//sweetDir := getSweetDir()
	//dayDir := t.Format("2006-01-02")

	//getList(sweetDir, dayDir)

	// TODO: Import any incomplete items from yesterday's list.
	// TODO: Try to open today's file and if it doesn't exist, create it.

	var list []listItem

	// TODO: Load the list
	item := listItem{status: "*", description: "Make the hop tool"}
	list = append(list, item)

	// Display the list
	displayList(list)

	// Prompt for command (add/remove/complete/quit)
	//cmd := make([]string, 0)
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
				scanner.Scan()
				item.description = scanner.Text()
				list = append(list, item)
				displayList(list)
			case "r":

				// Remove an item from the list
				fmt.Println("Removing item")
				fmt.Print("Item number > ")
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
				// TODO: Mark an item completed
				fmt.Println("Completing item")
			}
		}
		if cmd == "q" {
			break
		}
	}

	// TODO: Begin shutdown process

	// DEBUG
	fmt.Println("Shutting down...")
	// TODO: Save changes
	// TODO: Quit
}
