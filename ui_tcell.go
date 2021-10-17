// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

import (
	tb "github.com/nsf/termbox-go"
	"image"
)

var _ = &tcellUI{}

type tcellUI struct {
	painter *Painter
	root    Widget

	keybindings []*keybinding

	quit chan struct{}

	screen Screen

	kbFocus *kbFocusController

	eventQueue chan Event
}

func newTcellUI(root Widget) (*tcellUI, error) {
	p := NewPainter()
	root.SetPainter(p)
	return &tcellUI{
		painter:     p,
		root:        root,
		keybindings: make([]*keybinding, 0),
		quit:        make(chan struct{}, 1),
		kbFocus:     &kbFocusController{chain: DefaultFocusChain},
		eventQueue:  make(chan Event),
	}, nil
}

func (ui *tcellUI) Repaint() {
	ui.painter.Repaint(ui.root)
}

func (ui *tcellUI) SetWidget(w Widget) {
	ui.root = w
}

func (ui *tcellUI) SetFocusChain(chain FocusChain) {
	if ui.kbFocus.focusedWidget != nil {
		ui.kbFocus.focusedWidget.SetFocused(false)
	}

	ui.kbFocus.chain = chain
	ui.kbFocus.focusedWidget = chain.FocusDefault()

	if ui.kbFocus.focusedWidget != nil {
		ui.kbFocus.focusedWidget.SetFocused(true)
	}
}

func (ui *tcellUI) SetKeybinding(seq string, fn func()) {
	ui.keybindings = append(ui.keybindings, &keybinding{
		sequence: seq,
		handler:  fn,
	})
}

// ClearKeybindings reinitialises ui.keybindings so as to revert to a
// clear/original state
func (ui *tcellUI) ClearKeybindings() {
	ui.keybindings = make([]*keybinding, 0)
}

func (ui *tcellUI) Run() error {
	if err := tb.Init(); err != nil {
		return err
	}
	tb.SetInputMode(tb.InputEsc | tb.InputMouse)
	tb.SetOutputMode(tb.Output256)
	failed := true
	defer func() {
		if failed {
			tb.Close()
		}
	}()

	if w := ui.kbFocus.chain.FocusDefault(); w != nil {
		w.SetFocused(true)
		ui.kbFocus.focusedWidget = w
	}
	ui.screen.Clear()
	ui.reSize(nil)
	uiEvents := PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case KeyCtrlQ, "<C-c>":
				return nil
			default:
				ui.handleEvent(e)
			}
		//termui.Render(ui.root)
		case e := <-ui.eventQueue:
			ui.handleEvent(e)
		case w := <-ui.painter.drawQueue:
			ui.painter.Repaint(w)
		}
	}
}

func (ui *tcellUI) reSize(e *Event) {
	var w, h int
	if e == nil {
		w, h = tb.Size()
	} else {
		payload := e.Payload.(Resize)
		w, h = payload.Width, payload.Height
	}
	ui.root.Resize(image.Point{0, 0}, image.Pt(w, h))
	ui.Repaint()
}

func (ui *tcellUI) handleEvent(ev Event) {
	switch ev.Type {
	case KeyboardEvent:
		ui.root.DoEvent(ev)
	case MouseEvent:
		ui.root.DoEvent(ev)
	case ResizeEvent:
		ui.reSize(&ev)
		//ui.eventQueue <- paintEvent{}
	case CallbackEvent:
		// Gets stuck in a print loop when the logger is a widget.
		//logger.Printf("Received callback event")
	case PaintEvent:
		logger.Printf("Received paint event")
	}
}

// Quit signals to the UI to start shutting down.
func (ui *tcellUI) Quit() {
	logger.Printf("Quitting")
	tb.Close()
	ui.quit <- struct{}{}
}

// Schedule an update of the UI, running the given
// function in the UI goroutine.
//
// Use this to update the UI in response to external events,
// like a timer tick.
// This method should be used any time you call methods
// to change UI objects after the first call to `UI.Run()`.
//
// Changes invoked outside of either this callback or the
// other event handler callbacks may appear to work, but
// is likely a race condition.  (Run your program with
// `go run -race` or `go install -race` to detect this!)
//
// Calling Update from within an event handler, or from within an Update call,
// is an error, and will deadlock.

func (ui *tcellUI) Update(fn func()) {
	//blk := make(chan struct{})
	//ui.eventQueue <- callbackEvent{func() {
	//	fn()
	//	close(blk)
	//}}
	//<-blk
}
