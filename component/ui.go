// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package component

import (
	"github.com/gizak/termui/v3"
	"log"
	"sync"
)

// UI defines the operations needed by the underlying engine.
type UI interface {
	// SetWidget sets the root widget of the UI.
	//SetWidget(w Widget)
	// SetTheme sets the current theme of the UI.
	//SetTheme(p *Theme)
	// SetKeybinding sets the callback for when a key sequence is pressed.
	SetKeybinding(seq string, fn func())
	// ClearKeybindings removes all previous set keybindings.
	ClearKeybindings()
	// SetFocusChain sets a chain of widgets that determines focus order.
	//SetFocusChain(ch FocusChain)
	// Run starts the UI goroutine and blocks either Quit was called or an error occurred.
	Run() error
	// Update schedules work in the UI thread and await its completion.
	// Note that calling Update from the UI thread will result in deadlock.
	Update(fn func())
	// Quit shuts down the UI goroutine.
	Quit()
	// Repaint the UI
	Repaint()

	Close()
}

type ui struct {
	//widgets []Widget
	root Widget
	sync.RWMutex
}

func NewUi(root Widget) (UI, error) {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
		return nil, err
	}
	return &ui{root: root}, nil
}

func (s *ui) Close() {
	termui.Close()
}

func (s *ui) Run() error {
	termWidth, termHeight := termui.TerminalDimensions()
	s.TransportEvent(termui.Event{
		Type: termui.ResizeEvent,
		Payload: termui.Resize{
			Width:  termWidth,
			Height: termHeight,
		},
	})
	termui.Render(s.root)
	uiEvents := termui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case KeyCtrlQ, "<C-c>":
				return nil
			default:
				s.TransportEvent(e)
			}
			termui.Render(s.root)
		}
	}
	termui.Close()
	return nil
}

func (s *ui) Update(fn func()) {
	fn()
}

func (s *ui) Quit() {
}

func (s *ui) Repaint() {

}

func (s *ui) SetKeybinding(seq string, fn func()) {

}

func (s *ui) ClearKeybindings() {

}

func (s *ui) TransportEvent(e termui.Event) {
	s.root.DoEvent(e)
}
