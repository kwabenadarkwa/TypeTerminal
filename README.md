# TypeTerminal

This is a type test on the terminal that allows users to practice their typing.
It is built using Charm library things and it is based on the [Elm Architecture](https://guide.elm-lang.org/architecture/). This basically uses a _Model_, _View_ and _Update_

## Tools Used

- Charm Cli for the entire thing
- Wish for serving the terminal application through ssh
- Bubble Tea for building the ui

## How to Run

- In order to install things use,`go mod tidy`.
- In order to build things use,`go build main.go` or run with `go run .`.
- In order to run things use,`./main`.
- Starting ssh server `ssh -p 23234 localhost`

## TODO

- [x] Set things up
- [x] Figure out the high level architecture of shit
- [x] Create json file for quotes
- [x] figure out how to display text on terminal
- [x] figure out how to track text as the user types it
- [x] figure out how to track text that the user gets wrong
- [x] take care of the backspace case
- [x] track words that are right
- [x] figure out how to track WPM
- [x] figure out how to track the words that the user typed in right
- [x] Create a theme package
  - [x] Detect user's terminal theme in the beginning
- [x] figure out how to center things on the terminal
- [x] create the header and footer thing that displays things that will be constant
- [ ] create the screen for WPM and then logic that helps track the WPM
- [ ] start doing research into how it can actually be hosted and people can use it
- [ ] fix the ui issues on for when the terminal gets too small that the text is horrible. you should set like some size of the terminal that is minimum and that th euser to increase to to use it

## Refactoring

- [x] change the things that are supposed to be enums to enums
- [x] check on renaming of something that don't make sense
- [x] remove some of the functions that only call other functions
- [x] Create a large model to house all other models
