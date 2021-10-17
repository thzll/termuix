// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

import (
	"image"
	"strings"

	wordwrap "github.com/mitchellh/go-wordwrap"
)

var _ Widget = &Label{}

// Label is a widget to display read-only text.
type Label struct {
	Block
	text     string
	wordWrap bool

	// cache the result of SizeHint() (see #14)
	cacheSizeHint *image.Point

	styleName string
}

// NewLabel returns a new Label.
func NewLabel(text string) *Label {
	l := &Label{
		Block: *NewBlock(),
		text:  text,
	}
	l.style = NewStyle(ColorWhite)
	return l
}

// Resize changes the size of the Widget.
func (l *Label) Resize(pos, size image.Point) {
	if l.size != size {
		l.cacheSizeHint = nil
	}
	l.widgetBlock.SetRect(pos.X, pos.Y, pos.X+size.X, pos.Y+size.Y)
}

func (l *Label) draw() {
	l.Block.draw()
	p := l.GetPainter()
	if p == nil {
		return
	}
	lines := l.lines()
	for _, line := range lines {
		l.drawLine(line, p)
	}
}

// Draw draws the label.
func (l *Label) Draw() {
	l.Lock()
	defer l.Unlock()
	l.draw()
	l.drawSubWidget()
}

// Draw lines.
func (l *Label) drawLine(line string, p *Painter) {
	for {
		var ptext string
		inner := l.GetInnerRealPos()
		maxWidth := inner.Size().X - l.px
		if stringWidth(line) > maxWidth {
			ptext = line[:maxWidth]
			if l.wordWrap {
				line = line[maxWidth:]
			} else {
				line = ""
			}
		} else {
			ptext = line
			line = ""
		}
		if size := inner.Size(); l.px < size.X && l.py < size.Y {
			p.DrawText(inner.Min.X+l.px, inner.Min.Y+l.py, ptext, &l.style)
			l.py += 1
			l.px = 0
		}
		if len(line) == 0 {
			break
		}
	}
}

// MinSizeHint returns the minimum size the widget is allowed to be.
func (l *Label) MinSizeHint() image.Point {
	return image.Point{1, 1}
}

// SizeHint returns the recommended size for the label.
func (l *Label) SizeHint() image.Point {
	if l.cacheSizeHint != nil {
		return *l.cacheSizeHint
	}
	var max int
	lines := l.lines()
	for _, line := range lines {
		if w := stringWidth(line); w > max {
			max = w
		}
	}
	sizeHint := image.Point{max, len(lines)}
	l.cacheSizeHint = &sizeHint
	return sizeHint
}

func (l *Label) lines() []string {
	txt := l.text
	if l.wordWrap {
		txt = wordwrap.WrapString(l.text, uint(l.GetInner().Size().X))
		txt = wordwrap.WrapString(l.text, 5)
	}
	return strings.Split(txt, "\n")
}

// Text returns the text content of the label.
func (l *Label) Text() string {
	return l.text
}

// SetText sets the text content of the label.
func (l *Label) SetText(text string) {
	l.Lock()
	defer l.Unlock()
	l.cacheSizeHint = nil
	l.text = text
	l.rePaint(l)
}

// SetWordWrap sets whether text content should be wrapped.
func (l *Label) SetWordWrap(enabled bool) {
	l.wordWrap = enabled
}

// SetStyleName sets the identifier used for custom styling.
func (l *Label) SetStyleName(style string) {
	l.styleName = style
}
