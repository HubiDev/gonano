package gui

// ViewBase is the base for all views of the application
type ViewBase struct {
	Cur    *Cursor
	Header []string
	Footer []string
	Width  int
	Height int
}

// NewViewBase creates a new instance of the ViewBase struct
func NewViewBase() *ViewBase {
	return new(ViewBase)
}

// Draw
func (view *ViewBase) Draw() {

}
