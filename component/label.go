// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package component

import (
	. "github.com/gizak/termui/v3"
	"github.com/mitchellh/go-wordwrap"
	"image"
	"strings"
)

var _ = &Label{}

// Label is a widget to display read-only text.
type Label struct {
	WidgetBase
	wordWrap bool
	// cache the result of SizeHint() (see #14)
	cacheSizeHint *image.Point

	styleName string
}

// NewLabel returns a new Label.
func NewLabel(text string) *Label {
	label := &Label{
		WidgetBase: WidgetBase{
			text:        text,
			Block:       *NewBlock(),
			style:       NewStyle(ColorWhite, ColorBlue),
			sizePolicyX: Expanding,
			sizePolicyY: Expanding,
		},
	}
	label.SetRect(0, 0, 100, 20)
	return label
}

// Resize changes the size of the Widget.
func (l *Label) Resize(pos image.Point, size image.Point) {
	if l.Size() != size {
		l.cacheSizeHint = nil
	}
	l.WidgetBase.Resize(pos, size)
}

// Draw draws the label.
func (l *Label) Draw(buf *Buffer) {
	l.WidgetBase.Draw(buf)
	if l.text != "" {
		lines := l.lines()
		//buf.SetString(s.text, NewStyle(ColorWhite, ColorBlue),image.Pt(0, 0))
		inner := l.Inner
		if l.Border {
			inner = inner.Add(image.Point{1, 1})
		}
		for y, line := range lines {
			for x, rune := range line {
				buf.SetCell(
					NewCell(int32(rune), NewStyle(7)),
					image.Pt(inner.Min.X+x, inner.Min.Y+y-1),
				)
			}
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
		size := l.Size()
		txt = wordwrap.WrapString(l.text, uint(size.X))
	} else {
		size := l.Size()
		if l.Border {
			size.Sub(image.Point{10, 10})
		}
		txt = wordwrap.WrapString(l.text, 10)
		return []string{strings.Split(txt, "\n")[0]}
	}
	return strings.Split(txt, "\n")
}

// Text returns the text content of the label.
func (l *Label) Text() string {
	return l.text
}

// SetText sets the text content of the label.
func (l *Label) SetText(text string) {
	l.cacheSizeHint = nil
	l.text = text
}

// SetWordWrap sets whether text content should be wrapped.
func (l *Label) SetWordWrap(enabled bool) {
	l.wordWrap = enabled
}

// SetStyleName sets the identifier used for custom styling.
func (l *Label) SetStyleName(style string) {
	l.styleName = style
}
