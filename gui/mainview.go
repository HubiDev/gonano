package gui

import (
	"gonano/application"

	"github.com/nsf/termbox-go"
)

//MainView represnte the main view of the application
type MainView struct {
	*ViewBase
	KeyFuncs map[termbox.Key]func()
	Buffer   []string
}

// NewMainView created a new instance of the MainView struct
func NewMainView() (res *MainView) {

	res = new(MainView)
	res.ViewBase = NewViewBase()
	res.Cur = InitCursor()
	res.KeyFuncs = make(map[termbox.Key]func())

	res.KeyFuncs[termbox.KeyArrowLeft] = res.Cur.MoveLeft
	res.KeyFuncs[termbox.KeyArrowRight] = res.Cur.MoveRight
	res.KeyFuncs[termbox.KeyArrowUp] = res.Cur.MoveUp
	res.KeyFuncs[termbox.KeyArrowDown] = res.Cur.MoveDown
	res.KeyFuncs[termbox.KeyCtrlX] = application.QuitApp

	return
}

//Show displays the view --> blocks the current thread
func (view *MainView) Show() {

	//Clear content from potential view that was shown before
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlue)

	for {
		event := termbox.PollEvent()

		if event.Type == termbox.EventKey {
			view.HandleKeyEvent(event.Key)
		}
	}

}

// Draw prints all content from the buffer out to the terminal
func (view *MainView) Draw() {

	view.ViewBase.Draw()

	//TODO
}

// HandleKeyEvent handles a key press
func (view *MainView) HandleKeyEvent(key termbox.Key) {

	cmd := view.KeyFuncs[key]

	if cmd != nil {
		cmd()
	}

}
