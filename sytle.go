// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

// Color represents a color.
type Color int32

// Common colors.
const (
	ColorDefault Color = iota
	//ColorBlack
	//ColorWhite
	//ColorRed
	//ColorGreen
	//ColorBlue
	//ColorCyan
	//ColorMagenta
	//ColorYellow

	ColorBlack   Color = 0
	ColorRed     Color = 1
	ColorGreen   Color = 2
	ColorYellow  Color = 3
	ColorBlue    Color = 4
	ColorMagenta Color = 5
	ColorCyan    Color = 6
	ColorWhite   Color = 7
)

// Decoration represents a bold/underline/etc. state
type Decoration int

// Decoration modes: Inherit from parent widget, explicitly on, or explicitly off.
const (
	DecorationInherit Decoration = iota
	DecorationOn
	DecorationOff
)

const ColorClear Color = -1

type Modifier uint

const (
	// ModifierClear clears any modifiers
	ModifierClear     Modifier = 0
	ModifierBold      Modifier = 1 << 9
	ModifierUnderline Modifier = 1 << 10
	ModifierReverse   Modifier = 1 << 11
)

// Style determines how a cell should be painted.
// The zero value uses default from
type Style struct {
	Fg Color
	Bg Color

	Modifier Modifier
}

var StyleClear = Style{
	Fg:       ColorClear,
	Bg:       ColorClear,
	Modifier: ModifierClear,
}

// NewStyle takes 1 to 3 arguments
// 1st argument = Fg
// 2nd argument = optional Bg
// 3rd argument = optional Modifier
func NewStyle(fg Color, args ...interface{}) Style {
	bg := ColorClear
	modifier := ModifierClear
	if len(args) >= 1 {
		bg = args[0].(Color)
	}
	if len(args) == 2 {
		modifier = args[1].(Modifier)
	}
	return Style{
		fg,
		bg,
		modifier,
	}
}

func (s *Style) SetModifier(m Modifier) {
	s.Modifier = m
}

// mergeIn returns the receiver Style, with any changes in delta applied.
func (s Style) mergeIn(delta Style) Style {
	result := s
	if delta.Fg != ColorDefault {
		result.Fg = delta.Fg
	}
	if delta.Bg != ColorDefault {
		result.Bg = delta.Bg
	}
	return result
}
