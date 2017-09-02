package main

import (
	"github.com/nsf/termbox-go"
)

func main() {

	termbox.Init()
	defer termbox.Close() //clear stuff up after programm quits

}
