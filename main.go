package main

import (
	"gonano/gui"

	"github.com/nsf/termbox-go"
)

func main() {

	termbox.Init()
	defer termbox.Close() //clear stuff up after programm quits

	view := gui.NewMainView()
	view.Show()

}
