# Gofi

Go(ofy)-Rofi

[Rofi](https://github.com/davatorium/rofi) is an application launcher, this program is designed to emulate the [dmenu](https://tools.suckless.org/dmenu/) function of Rofi.  
I made this as I was trying to find a Rofi/dmenu alternative that works on Wayland and didn't require a bunch of dependancies, and this was proving difficult.

Built using [Go-FLTK](https://github.com/pwiecz/go-fltk) as the UI library

Still heavily WIP.  
ToDo:
* Add title/prompt
* Add pages
* Add ability to display icons
* Allow program to read from stdin instead of just a specified file
* Remove the window borders (like Rofi or dmenu)
* Ability to define custom hotkeys

## Installation

`go mod tidy`  
`go build gofi.go`

## Usage

`./gofi input.txt` for displaying a list of options  
`./gofi -i` for getting user input and returning it

Pressing `Return` on an entry will print that entry, which can then be used by another program.  
`Shift + Return` will print the entry and exit with code 12  
`Ctrl + w` will print the entry and return with code 11  
`Ctrl + a` will return with code 10

This can be used to define different actions in the calling programs, based on the return code. For example a simple bookmarking script. `Return` can be used to copy the entry to the clipboard, `Shift + Return` to open it in a web browser, `Ctrl + w` to remove an entry and `Ctrl + a` to add a new entry.
