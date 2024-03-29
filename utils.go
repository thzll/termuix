// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

import (
	"github.com/mattn/go-runewidth"
	"unicode/utf8"
)

// runeWidth returns the cell width of given rune
func runeWidth(r rune) int {
	return runewidth.RuneWidth(r)
}

// stringWidth returns the cell width of given string
func stringWidth(s string) int {
	return runewidth.StringWidth(s)
}

// trimRightLen returns s with n runes trimmed off
func trimRightLen(s string, n int) string {
	if n <= 0 {
		return s
	}
	c := utf8.RuneCountInString(s)
	runeCount := 0
	var i int
	for i = range s {
		if runeCount >= c-n {
			break
		}
		runeCount++
	}
	return s[:i]
}
