// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

import (
	rw "github.com/mattn/go-runewidth"
	"image"
)

// Surface defines a surface that can be painted on.
type Surface interface {
	GetCell(p image.Point) Cell
	SetCell(c Cell, p image.Point)
	Fill(c Cell, rect image.Rectangle)
	SetCursor(x, y int)
	HideCursor()
	Size() image.Point
	Clear()
	Show()
}

// Painter provides operations to paint on a surface.
type Painter struct {
	// Surface to paint on.
	surface Surface
	// Transform stack
	transforms []image.Point
	drawQueue  chan Widget
}

// NewPainter returns a new instance of Painter.
func NewPainter() *Painter {
	return &Painter{
		surface:   NewScreen(image.Rectangle{}),
		drawQueue: make(chan Widget, 100),
	}
}

// Translate pushes a new translation transform to the stack.
func (p *Painter) Translate(x, y int) {
	p.transforms = append(p.transforms, image.Point{x, y})
}

// Restore pops the latest transform from the stack.
func (p *Painter) Restore() {
	if len(p.transforms) > 0 {
		p.transforms = p.transforms[:len(p.transforms)-1]
	}
}

// Begin prepares the surface for painting.
func (p *Painter) Begin() {
	p.surface.Clear()
}

// End finalizes any painting that has been made.
func (p *Painter) End() {
	p.surface.Show()
}

func (p *Painter) addPaint(w Widget) {
	p.drawQueue <- w
}

// Repaint clears the surface, draws the scene and flushes it.
func (p *Painter) Repaint(w Widget) {
	p.surface.HideCursor()
	p.Begin()
	w.Draw()
	p.End()
}

// DrawCursor draws the cursor at the given position.
func (p *Painter) DrawCursor(x, y int) {
	p.surface.SetCursor(x, y)
}

// DrawRune paints a rune at the given coordinate.
func (p *Painter) DrawRune(x, y int, r rune, st *Style) {
	p.SetCell(Cell{r, *st}, image.Pt(x, y))
}

// DrawText paints a string starting at the given coordinate.
func (p *Painter) DrawText(x, y int, text string, st *Style) {
	for _, r := range text {
		p.DrawRune(x, y, r, st)
		x += runeWidth(r)
	}
}

func (self *Painter) GetCell(p image.Point) Cell {
	return self.surface.GetCell(p)
}

func (self *Painter) SetCell(c Cell, p image.Point) {
	self.surface.SetCell(c, p)
}

func (self *Painter) Fill(c Cell, rect image.Rectangle) {
	self.surface.Fill(c, rect)
}

func (self *Painter) SetString(s string, style Style, p image.Point) {
	runes := []rune(s)
	x := 0
	for _, char := range runes {
		self.SetCell(Cell{char, style}, image.Pt(p.X+x, p.Y))
		x += rw.RuneWidth(char)
	}
}
