# TypeTerminal

This is a type test on the terminal that allows users to practice their typing.

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
- [ ] figure out how to track WPM
- [ ] Create a theme package 
    - [ ] Detect user's terminal theme in the beginning 

## UI

- [ ] figure out how to center things on the terminal

## Refactoring

- [x] change the things that are supposed to be enums to enums
- [ ] check on renaming of something that don't make sense
- [ ] Create a large model to house all other models (Reference our lord and saviour Primeagen)

