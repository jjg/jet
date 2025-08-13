# jet
Journal EnTry

A reliable focused writing tool.

![screenshot of jet running in a terminal](docs/images/screenshot.png)

## Status
Jet is close to complete.  This project is evolving into something a bit broader than just Jet, so expect to see more changes soon.

## Installation
Download the `jet` binaries from the [bin](bin) directory for your architecture/os and put it somewhere in your path.

Alternatively (if you have Go installed), install it with `go install ./cmd/jet/`, `go install ./cmd/today/`, etc..

## Usage
1. Open a terminal and type `jet`
2. Type your journal entry
3. When you're done, press return on a blank line

`jet` will create a directory called `jet-journal` in your home directory where it will keep a new journal file for each day containing the entries you record as shown above.  Feel free to edit and consume these files however you like, `jet` will simply append to the current day's file if a new entry is recorded for that day.

`jet help`: displays some info and instructions.

To view journal entries, use these additional tools:
* `today`: displays the contents of the current day's journal entries
* `yesterday`: displays the contents of yesterday's entries

## TODO:
Some potential ideas for the future.

- [X] `today` subcommand
- [X] `yesterday` subcommand
- [ ] `tomorrow` subcommand
- [ ] `week` subcommand
- [X] `help` subcommand
- [ ] `find` subcommand
- [X] Add linefeeds between entries
- [ ] Provide better errors (don't just `panic()` all the time)
- [ ] Allow journal dir to be customized?
- [X] Allow text to be piped-in?
- [X] Add a build script so binaries can be built in one step
- [ ] Add date to header of each journal file?
- [X] Tweak journal file permissions (only accessible to owner by default)
- [ ] Arbitrary date subcommand?
- [ ] Jet for lists?

## Contributing
This section is at the end because right now it's not intended for anyone other than the author.

### Guidelines
* No developers, no users: the programs in here are tools, the people who make them toolmakers, the people who employ them authors, accountants, filmmakers, etc.
* Each tool is implemented as a single, stand-alone source code file (if it gets too complicated for one file, it should be broken down into separate tools, not source files)
* Dependencies will be heavily scrutinized; tools should only depend on other tools not modules, shared source files, etc.
* Changes (PR's, etc.) should be small and simple enough to complete in a day, ideally in a single work session
