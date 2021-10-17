// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

import (
	tb "github.com/nsf/termbox-go"
	"image"
	"sync"

	rw "github.com/mattn/go-runewidth"
)

// Cell represents a viewable terminal cell
type Cell struct {
	Rune  rune
	Style Style
}

var CellClear = Cell{
	Rune:  ' ',
	Style: StyleClear,
}

// NewCell takes 1 to 2 arguments
// 1st argument = rune
// 2nd argument = optional style
func NewCell(rune rune, args ...interface{}) Cell {
	style := StyleClear
	if len(args) == 1 {
		style = args[0].(Style)
	}
	return Cell{
		Rune:  rune,
		Style: style,
	}
}

// Buffer represents a section of a terminal and is a renderable rectangle of cells.
type Screen struct {
	image.Rectangle
	CellMap map[image.Point]Cell
	sync.Mutex
}

func NewScreen(r image.Rectangle) *Screen {
	buf := &Screen{
		Rectangle: r,
		CellMap:   make(map[image.Point]Cell),
	}
	buf.Fill(CellClear, r) // clears out area
	return buf
}

func (self *Screen) GetCell(p image.Point) Cell {
	return self.CellMap[p]
}

func (self *Screen) SetCell(c Cell, p image.Point) {
	if self.CellMap == nil {
		self.CellMap = make(map[image.Point]Cell)
	}
	self.CellMap[p] = c
}

func (self *Screen) Fill(c Cell, rect image.Rectangle) {
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			self.SetCell(c, image.Pt(x, y))
		}
	}
}

func (self *Screen) Size() image.Point {
	tb.Sync()
	width, height := tb.Size()
	return image.Point{width, height}
}

func (self *Screen) SetString(s string, style Style, p image.Point) {
	runes := []rune(s)
	x := 0
	for _, char := range runes {
		self.SetCell(Cell{char, style}, image.Pt(p.X+x, p.Y))
		x += rw.RuneWidth(char)
	}
}

func (s *Screen) SetCursor(x, y int) {

}

func (s *Screen) HideCursor() {

}

func (s *Screen) Clear() {
	tb.Clear(tb.ColorDefault, tb.Attribute(ColorDefault))
}

func (s *Screen) Show() {
	for point, cell := range s.CellMap {
		tb.SetCell(
			point.X, point.Y,
			cell.Rune,
			tb.Attribute(cell.Style.Fg+1)|tb.Attribute(cell.Style.Modifier), tb.Attribute(cell.Style.Bg+1),
		)
	}
	tb.Flush()
}
