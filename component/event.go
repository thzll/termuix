// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package component

const (
	KeyF1         = "<F1>"
	KeyF2         = "<F2>"
	KeyF3         = "<F3>"
	KeyF4         = "<F4>"
	KeyF5         = "<F5>"
	KeyF6         = "<F6>"
	KeyF7         = "<F7>"
	KeyF8         = "<F8>"
	KeyF9         = "<F9>"
	KeyF10        = "<F10>"
	KeyF11        = "<F11>"
	KeyF12        = "<F12>"
	KeyInsert     = "<Insert>"
	KeyDelete     = "<Delete>"
	KeyHome       = "<Home>"
	KeyEnd        = "<End>"
	KeyPgup       = "<PageUp>"
	KeyPgdn       = "<PageDown>"
	KeyArrowUp    = "<Up>"
	KeyArrowDown  = "<Down>"
	KeyArrowLeft  = "<Left>"
	KeyArrowRight = "<Right>"

	KeyCtrlSpace  = "<C-<Space>>" //  KeyCtrl2  KeyCtrlTilde
	KeyCtrlA      = "<C-a>"
	KeyCtrlB      = "<C-b>"
	KeyCtrlC      = "<C-c>"
	KeyCtrlD      = "<C-d>"
	KeyCtrlE      = "<C-e>"
	KeyCtrlF      = "<C-f>"
	KeyCtrlG      = "<C-g>"
	KeyBackspace  = "<C-<Backspace>>" //  KeyCtrlH
	KeyTab        = "<Tab>"           //  KeyCtrlI
	KeyCtrlJ      = "<C-j>"
	KeyCtrlK      = "<C-k>"
	KeyCtrlL      = "<C-l>"
	KeyEnter      = "<Enter>" //  KeyCtrlM
	KeyCtrlN      = "<C-n>"
	KeyCtrlO      = "<C-o>"
	KeyCtrlP      = "<C-p>"
	KeyCtrlQ      = "<C-q>"
	KeyCtrlR      = "<C-r>"
	KeyCtrlS      = "<C-s>"
	KeyCtrlT      = "<C-t>"
	KeyCtrlU      = "<C-u>"
	KeyCtrlV      = "<C-v>"
	KeyCtrlW      = "<C-w>"
	KeyCtrlX      = "<C-x>"
	KeyCtrlY      = "<C-y>"
	KeyCtrlZ      = "<C-z>"
	KeyEsc        = "<Escape>" //  KeyCtrlLsqBracket  KeyCtrl3
	KeyCtrl4      = "<C-4>"    //  KeyCtrlBackslash
	KeyCtrl5      = "<C-5>"    //  KeyCtrlRsqBracket
	KeyCtrl6      = "<C-6>"
	KeyCtrl7      = "<C-7>" //  KeyCtrlSlash  KeyCtrlUnderscore
	KeySpace      = "<Space>"
	KeyBackspace2 = "<Backspace>" //  KeyCtrl8:
)

func isCharKey(keyId string) bool {
	if len(keyId) > 0 && keyId[0] != '<' {
		return true
	}
	return false
}
