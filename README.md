# Gofi

Go(ofy)-Rofi

[Rofi](https://github.com/davatorium/rofi) is an application launcher, this program is designed to emulate the [dmenu](https://tools.suckless.org/dmenu/) function of Rofi.  
I made this as trying to find a Rofi/dmenu alternative that works on Wayland and doesn't require a bunch of dependancies was proving difficult. 

Built using [Go-FLTK](https://github.com/pwiecz/go-fltk) as the UI library

Still heavily WIP.  
ToDo:
* Add title/prompt
* Add pages
* Add ability to display icons
* Allow program to read from stdin instead of just a specified file
* Remove the window borders (like Rofi or dmenu)

## Installation

`go mod tidy`
`go build gofi.go`

## Usage

`./gofi input.txt` for displaying a list of options  
`./gofi -i` for getting user input and returning it
