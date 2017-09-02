package application

import (
	"log"

	termbox "github.com/nsf/termbox-go"
)

// QuitApp quits the application
func QuitApp() {
	termbox.SetCursor(0, 0)
	log.Fatal("Application was stopped by user.")
}
