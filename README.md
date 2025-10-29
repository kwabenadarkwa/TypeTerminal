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
- [ ] figure out how to track WPM
- [ ] figure out how to track the words that the user typed in right
- [ ] Create a theme package
  - [x] Detect user's terminal theme in the beginning
  - [ ] create a component for creating borders around things in the theme file
- [ ] things

## UI

- [ ] figure out how to center things on the terminal

## Refactoring

- [x] change the things that are supposed to be enums to enums
- [ ] check on renaming of something that don't make sense
- [ ] remove some of the functions that only call other functions
- [ ] Create a large model to house all other models
