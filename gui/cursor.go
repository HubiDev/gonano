package gui

import (
	"github.com/nsf/termbox-go"
)

// Cursor represents the cursor ind the console
type Cursor struct {
	PosX   int
	PosY   int
	LimitX int
	LimitY int
}

// InitCursor creates and initializes a new instance of the Cursor struct
func InitCursor() *Cursor {

	res := new(Cursor)
	res.LimitX, res.LimitY = termbox.Size()
	res.PosX = 0
	res.PosY = 0

	return res
}

// MoveLeft moves the cursor left
func (cur *Cursor) MoveLeft() {

	newPos := cur.PosX - 1

	if newPos >= 0 {
		cur.PosX = newPos
	}

	cur.UpdatePos()
}

// MoveRight moves the cursor right
func (cur *Cursor) MoveRight() {

	newPos := cur.PosX + 1

	if newPos < cur.LimitX {
		cur.PosX = newPos
	}

	cur.UpdatePos()
}

// MoveUp moves the cursor up
func (cur *Cursor) MoveUp() {

	newPos := cur.PosY - 1

	if newPos >= 0 {
		cur.PosY = newPos
	}

	cur.UpdatePos()

}

// MoveDown moves the cursor down
func (cur *Cursor) MoveDown() {

	newPos := cur.PosY + 1

	if newPos < cur.LimitY {
		cur.PosY = newPos
	}

	cur.UpdatePos()

}

// ResetPos sets the position of the cursor to 0,0
func (cur *Cursor) ResetPos() {

	cur.PosX = 0
	cur.PosY = 0
	cur.UpdatePos()
}

// UpdatePos sets the position of the cursor to the given coordinates
func (cur *Cursor) UpdatePos() {
	termbox.SetCursor(cur.PosX, cur.PosY)
}
