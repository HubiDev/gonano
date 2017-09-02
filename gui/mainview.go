package gui

type MainView struct {
	Cur *Cursor
}

// NewMainView created a new instance of the MainView struct
func NewMainView() (res *MainView) {

	res = new(MainView)
	res.Cur = InitCursor()

	return
}

func Run() {

}

func Draw() {

}
