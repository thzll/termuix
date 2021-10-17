// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package component

import (
	. "github.com/gizak/termui/v3"
	tb "github.com/nsf/termbox-go"
	"image"
)

// EchoMode is used to determine the visibility of Input text.
type EchoMode int

const (
	// EchoModeNormal displays the characters as they're being entered.
	EchoModeNormal EchoMode = iota

	// EchoModeNoEcho disables text display. This is useful for when the length
	// of the password should be kept secret.
	EchoModeNoEcho

	// EchoModePassword replaces all characters with asterisks.
	EchoModePassword
)

// Input is a one-line text editor. It lets the user supply the application
// with text, e.g., to input user and password information.
type Input struct {
	WidgetBase
	text RuneBuffer

	onTextChange func(*Input)
	onSubmit     func(*Input)

	echoMode EchoMode
	offset   int
}

// NewInput returns a new Input.
func NewInput() *Input {
	input := &Input{
		WidgetBase: WidgetBase{
			Block:       *NewBlock(),
			style:       NewStyle(ColorWhite, ColorBlue),
			sizePolicyY: Minimum,
			size:        image.Point{0, 1},
		},
	}
	input.SetRect(0, 0, 100, 50)
	input.SetFocused(true)
	return input
}

// Resize changes the size of the Widget.
func (l *Input) Resize(pos image.Point, size image.Point) {
	l.WidgetBase.Resize(pos, size)
}

// Draw draws the Input.
func (e *Input) Draw(buf *Buffer) {
	e.WidgetBase.Draw(buf)
	text := e.visibleText()
	buf.SetString(
		text,
		NewStyle(ColorWhite),
		image.Pt(
			e.Inner.Min.X,
			e.Inner.Min.Y,
		),
	)
	if e.IsFocused() {
		var off int
		if e.echoMode != EchoModeNoEcho {
			off = e.text.CursorPos().X - e.offset
		}

		tb.SetCursor(e.Inner.Min.X+off, e.Inner.Min.Y)
	}
}

// SizeHint returns the recommended size hint for the Input.
func (e *Input) SizeHint() image.Point {
	return image.Point{10, 1}
}

func (s *Input) DoEvent(ev Event) {
	switch ev.Type {
	case KeyboardEvent:
		s.DoKeyEvent(ev)
	case MouseEvent:
	case ResizeEvent:
	default:

	}

}

// OnKeyEvent handles key events.
func (e *Input) DoKeyEvent(ev Event) {
	if !e.IsFocused() {
		return
	}
	screenWidth := e.Size().X
	e.text.SetMaxWidth(screenWidth)

	if !isCharKey(ev.ID) {
		switch ev.ID {
		case KeyEnter:
			if e.onSubmit != nil {
				e.onSubmit(e)
			}
		case KeyBackspace:
			fallthrough
		case KeyBackspace2:
			e.text.Backspace()
			if e.offset > 0 && !e.isTextRemaining() {
				e.offset--
			}
			if e.onTextChange != nil {
				e.onTextChange(e)
			}
		case KeyDelete, KeyCtrlD:
			e.text.Delete()
			if e.onTextChange != nil {
				e.onTextChange(e)
			}
		case KeyArrowLeft, KeyCtrlB:
			e.text.MoveBackward()
			if e.offset > 0 {
				e.offset--
			}
		case KeyArrowRight, KeyCtrlF:
			e.text.MoveForward()

			isCursorTooFar := e.text.CursorPos().X >= screenWidth
			isTextLeft := (e.text.Width() - e.offset) > (screenWidth - 1)

			if isCursorTooFar && isTextLeft {
				e.offset++
			}
		case KeyHome, KeyCtrlA:
			e.text.MoveToLineStart()
			e.offset = 0
		case KeyEnd, KeyCtrlE:
			e.text.MoveToLineEnd()
			e.ensureCursorIsVisible()
		case KeyCtrlK:
			e.text.Kill()
		}
		return
	}
	idCode := []rune(ev.ID)[0]
	e.text.WriteRune(rune(idCode))
	if e.text.CursorPos().X >= screenWidth {
		e.offset++
	}
	if e.onTextChange != nil {
		e.onTextChange(e)
	}
}

// OnChanged sets a function to be run whenever the content of the Input has
// been changed.
func (e *Input) OnChanged(fn func(Input *Input)) {
	e.onTextChange = fn
}

// OnSubmit sets a function to be run whenever the user submits the Input (by
// pressing KeyEnter).
func (e *Input) OnSubmit(fn func(Input *Input)) {
	e.onSubmit = fn
}

// SetEchoMode sets the echo mode of the Input.
func (e *Input) SetEchoMode(m EchoMode) {
	e.echoMode = m
}

// SetText sets the text content of the Input.
func (e *Input) SetText(text string) {
	e.text.Set([]rune(text))
	// TODO: Enable when RuneBuf supports cursor movement for CJK.
	// e.ensureCursorIsVisible()
	e.offset = 0
}

func (e *Input) ensureCursorIsVisible() {
	left := e.text.Width() - (e.Size().X - 1)
	if left >= 0 {
		e.offset = left
	} else {
		e.offset = 0
	}
}

// Text returns the text content of the Input.
func (e *Input) Text() string {
	return e.text.String()
}

func (e *Input) visibleText() string {
	text := e.text
	if text.Len() == 0 {
		return ""
	}
	windowStart := e.offset
	windowEnd := e.Size().X + windowStart
	if windowEnd > text.Len() {
		windowEnd = text.Len()
	}
	return string(text.Runes()[windowStart:windowEnd])
}

func (e *Input) isTextRemaining() bool {
	return e.text.Width()-e.offset > e.Size().X
}
