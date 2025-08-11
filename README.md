# jet
Journal EnTry

A reliable focused writing tool.

## Status
Good enough.  See the TODO section below.

## Installation
Download the `jet` binary for your architecture/os and put it somewhere in your path.

Alternatively (if you have Go installed), install it with `go install`.

### Other platforms
Get a list of operating systems and architectures available.
```
go tool dist list
```

...then build for each platform (substitute `os` and `arch` in the filename appropriately):
```
GOOS=linux GOARCH=amd64 go build -o bin/jet-os-arch
```

## Usage
1. Open a terminal and type `jet`
2. Type your journal entry
3. When you're done, press return on a blank line

`jet` will create a directory called `jet-journal` in your home directory where it will keep a new journal file for each day containing the entries you record as shown above.  Feel free to edit and consume these files however you like, `jet` will simply append to the current day's file if a new entry is recorded for that day.

### Subcommands
Right now there is only one: `today`.  Running `jet today` will display the contents of the current day's journal entries.

...more to come.

## TODO
Some potential ideas for the future.

- [ ] `yesterday` subcommand
- [ ] `week` subcommand
- [ ] `help` subcommand
- [ ] `find` subcommand
- [ ] Add timestamps & linefeeds between entries
- [ ] Provide better errors (don't just panic() all the time).
- [ ] Allow journal dir to be customized?
- [ ] Allow text to be piped-in?
- [ ] Add a build script so binaries can be built in one step
