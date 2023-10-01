package main

import (
	"fmt"
	"os"
	"github.com/pwiecz/go-fltk"
	"strings"
)

const BACKGROUND=0x22212D00
const SELECTED=0x44475A00
const BLUE = 0x42A5F500
const GREY = 0x2B2B2B00

var curString string
var curSelected int
var curPage int
var bars [15]*fltk.Box
var lines []string

func resetCursor(oldSel int) {
	bars[oldSel].SetColor(BACKGROUND)
	bars[oldSel].Redraw()
	bars[0].SetColor(SELECTED)
	bars[0].Redraw()
}

func filterList() []string {
	toReturn := []string{}
	for _, entry := range lines {
		if strings.Contains(strings.ToLower(entry), strings.ToLower(curString)) {
			toReturn = append(toReturn, entry)
		}
	}
	return toReturn
}

func fillBars(toUse []string) {
	var curPageLen int = len(toUse) - curPage*15
	for i, _ := range bars {
		if i < curPageLen {
			bars[i].SetLabel(toUse[i+curPage*15])
			bars[i].Redraw()
		} else {
			bars[i].SetLabel("")
			bars[i].Redraw()
		}
	}
}

func getInput() {
	win := fltk.NewWindow(700, 700)
	win.SetColor(BACKGROUND)
	inputBar := fltk.NewBox(fltk.FLAT_BOX, 0, 0, 700, 60, curString)
	inputBar.SetColor(BACKGROUND)
	inputBar.SetLabelSize(22)
	inputBar.SetLabelColor(fltk.WHITE)
	inputBar.SetLabelFont(10)
	
	win.SetEventHandler(func(event fltk.Event) bool {
	    switch fltk.EventType() {
	    case fltk.KEY:
	    	if fltk.EventKey() == fltk.ENTER_KEY {
	    		fmt.Println(curString)
	    		if fltk.EventState() == 65536 {
	    			os.Exit(12) // Will show as 1 when using go run, should be fine when built with go build
	    		} else {
	    			os.Exit(0)
	    		}
	    	} else if fltk.EventKey() == fltk.BACKSPACE {
	    		curString = curString[:len(curString)-1]
	    		inputBar.SetLabel(curString)
	    	} else if fltk.EventKey() != 65505 && fltk.EventKey() != 65507 {
	    		curString = curString + string(fltk.EventKey())
	    		inputBar.SetLabel(curString)
	    	}
        case fltk.CLOSE:
            win.Destroy()
        }
        return false
    })
    
    win.End()
	win.Show()
	fltk.Run()
}

func main() {
	curSelected = 0
	curPage = 0
	curString = ""
	
	if len(os.Args) != 2 {
	    fmt.Println("Incorrect arguments")
	    os.Exit(1)
	}
	
	fltk.SetFont(10, "Roboto Regular")
	fltk.SetFont(11, "Roboto Regular Bold")
	
	if os.Args[1] == "-i" {
	    getInput()
        os.Exit(0)
	}
	
	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	lines = strings.Split(string(dat), "\n")
	//fmt.Println(lines)
	
	
	//fmt.Println("HERE")
	win := fltk.NewWindow(700, 700)
	win.SetPosition(610, 190)
	win.SetColor(BACKGROUND)
	inputBar := fltk.NewBox(fltk.FLAT_BOX, 0, 0, 700, 60, curString)
	inputBar.SetColor(BACKGROUND)
	inputBar.SetLabelSize(22)
	inputBar.SetLabelColor(fltk.WHITE)
	inputBar.SetLabelFont(10)
	
	for i, _ := range bars {
	    if i < len(lines) {
		    bars[i] = fltk.NewBox(fltk.FLAT_BOX, 0, i*40+65, 700, 30, lines[i+curPage*15])
		} else {
		    bars[i] = fltk.NewBox(fltk.FLAT_BOX, 0, i*40+65, 700, 30, "")
		}
		bars[i].SetColor(BACKGROUND)
		bars[i].SetLabelColor(fltk.WHITE)
	    bars[i].SetLabelFont(10)
	    bars[i].SetLabelSize(18)
	}
	bars[curSelected].SetColor(SELECTED)
	
	win.SetEventHandler(func(event fltk.Event) bool {
	    switch fltk.EventType() {
	    case fltk.KEY:
	    	if fltk.EventKey() == fltk.UP {
	    		curSelected--
	    		bars[curSelected+1].SetColor(BACKGROUND)
	    		bars[curSelected+1].Redraw()
	    		bars[curSelected].SetColor(SELECTED)
	    		bars[curSelected].Redraw()
	    	} else if fltk.EventKey() == fltk.DOWN {
	    		curSelected++
	    		bars[curSelected-1].SetColor(BACKGROUND)
	    		bars[curSelected-1].Redraw()
	    		bars[curSelected].SetColor(SELECTED)
	    		bars[curSelected].Redraw()
	    	} else if fltk.EventKey() == fltk.ENTER_KEY {
	    		fmt.Println(filterList()[curSelected+curPage*15])
	    		if fltk.EventState() == 65536 {
	    			os.Exit(12) // Will show as 1 when using go run, should be fine when built with go build
	    		} else {
	    			os.Exit(0)
	    		}
	    	} else if fltk.EventKey() == fltk.BACKSPACE {
	    		curString = curString[:len(curString)-1]
	    		inputBar.SetLabel(curString)
	    		resetCursor(curSelected)
	    		curPage = 0
	    		curSelected = 0
	    		fillBars(filterList())
	    	} else if fltk.EventKey() != 65505 && fltk.EventKey() != 65507 {
	    	    if fltk.EventState() == 262144 {
	    	        if fltk.EventKey() == 'w' {
	    	            fmt.Println(filterList()[curSelected+curPage*15])
	    	            os.Exit(11)
	    	        } else if fltk.EventKey() == 'a' {
	    	            os.Exit(10)
	    	        }
	    	    }
	    	//fmt.Println(fltk.EventState())
	    		curString = curString + string(fltk.EventKey())
	    		inputBar.SetLabel(curString)
	    		resetCursor(curSelected)
	    		curPage = 0
	    		curSelected = 0
	    		fillBars(filterList())
	    	}
        case fltk.CLOSE:
            win.Destroy()
        }
        return false
    })
	
	
	win.End()
	win.Show()
	fltk.Run()
}
