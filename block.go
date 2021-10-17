// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

import (
	"image"
)

// Block is the base struct inherited by most widgets.
// Block manages size, position, border, and title.
// It implements all 3 of the methods needed for the `Drawable` interface.
// Custom widgets will override the Draw method.
type Block struct {
	WidgetBase
}

func NewBlock() *Block {
	return &Block{
		WidgetBase: WidgetBase{
			widgetBlock: widgetBlock{
				Border:       true,
				BorderStyle:  Theme.Block.Border,
				BorderLeft:   true,
				BorderRight:  true,
				BorderTop:    true,
				BorderBottom: true,
				TitleStyle:   Theme.Block.Title,
			},
			layout:      Horizontal,
			sizePolicyX: Expanding,
			sizePolicyY: Expanding,
			style:       NewStyle(ColorWhite, ColorClear),
		},
	}
}

func (s *Block) drawBorder(p *Painter) {
	verticalCell := Cell{VERTICAL_LINE, s.BorderStyle}
	horizontalCell := Cell{HORIZONTAL_LINE, s.BorderStyle}
	min := s.GetParentMin()
	min.X += int(s.MarginLeft + s.X)
	min.Y += int(s.MarginTop + s.Y)
	max := image.Pt(min.X+int(s.Width)-s.MarginRight, min.Y+int(s.Height)-s.MarginBottom)
	// draw lines
	if s.BorderTop {
		p.Fill(horizontalCell, image.Rect(min.X+1, min.Y, max.X-1, min.Y+1))
	}
	if s.BorderBottom {
		p.Fill(horizontalCell, image.Rect(min.X+1, max.Y-1, max.X-1, max.Y))
	}
	if s.BorderLeft {
		p.Fill(verticalCell, image.Rect(min.X, min.Y+1, min.X+1, max.Y-1))
	}
	if s.BorderRight {
		p.Fill(verticalCell, image.Rect(max.X-1, min.Y+1, max.X, max.Y-1))
	}

	// draw corners
	if s.BorderTop && s.BorderLeft {
		p.SetCell(Cell{TOP_LEFT, s.BorderStyle}, min)
	}
	if s.BorderTop && s.BorderRight {
		p.SetCell(Cell{TOP_RIGHT, s.BorderStyle}, image.Pt(min.X+int(s.Width)-1, min.Y))
	}
	if s.BorderBottom && s.BorderLeft {
		p.SetCell(Cell{BOTTOM_LEFT, s.BorderStyle}, image.Pt(min.X, min.Y+int(s.Height)-1))
	}
	if s.BorderBottom && s.BorderRight {
		p.SetCell(Cell{BOTTOM_RIGHT, s.BorderStyle}, image.Pt(max.X-1, max.Y-1))
	}
}

func (s *Block) draw() {
	s.WidgetBase.draw()
	p := s.GetPainter()
	if p == nil {
		return
	}
	if s.Border {
		s.drawBorder(p)
	}
	if s.Title != "" {
		min := s.GetParentMin()
		min.X += int(s.MarginLeft) + 1
		min.Y += int(s.MarginTop)
		p.SetString(
			s.Title,
			s.TitleStyle,
			min,
		)
	}
}

// Draw implements the Drawable interface.
func (s *Block) Draw() {
	s.Lock()
	defer s.Unlock()
	s.draw()
	s.drawSubWidget()
}

// SetBorder sets whether the border is visible or not.
func (s *Block) SetBorder(enabled bool) {
	s.Border = enabled
}

// SetTitle sets the title of the box.
func (s *Block) SetTitle(title string) {
	s.Title = title
}

// GetRect implements the Drawable interface.
func (s *Block) GetOuter() image.Rectangle {
	return s.WidgetBase.GetOuter()
}

func (s *Block) Resize(pos image.Point, size image.Point) {
	s.WidgetBase.Resize(pos, size)
}
