// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

import "image"

type widgetBlock struct {
	rect                                                 image.Rectangle //
	Border                                               bool
	BorderStyle                                          Style
	X, Y                                                 int
	Width, Height                                        int
	BorderLeft, BorderRight, BorderTop, BorderBottom     bool
	PaddingLeft, PaddingRight, PaddingTop, PaddingBottom int
	MarginLeft, MarginRight, MarginTop, MarginBottom     int
	Title                                                string
	TitleStyle                                           Style
}

// GetRect implements the Drawable interface.
func (s *widgetBlock) GetOuter() image.Rectangle {
	return image.Rect(s.X, s.Y, s.X+s.Width, s.Y+s.Height)
	//w, h := s.Width + s.X, s.Height + s.Y
	//if s.Border {
	//	if s.BorderLeft {
	//		w += 1
	//	}
	//	if s.BorderTop {
	//		h += 1
	//	}
	//	if s.BorderRight {
	//		w += 1
	//	}
	//	if s.BorderBottom {
	//		h += 1
	//	}
	//}
	//w = w + s.MarginLeft + s.MarginRight
	//h = h + s.MarginTop + s.MarginBottom
	//return image.Rect(0, 0, int(w), int(h))
}

// GetRect GetInner the Drawable interface.
func (s *widgetBlock) GetInner() image.Rectangle {
	x, y := s.MarginLeft+s.X, s.MarginTop+s.Y
	w := int(s.Width)
	h := int(s.Height)
	if s.Border {
		if s.BorderLeft {
			x += 1
			w -= 1
		}
		if s.BorderTop {
			y += 1
			h -= 1
		}
		if s.BorderRight {
			w -= 1
		}
		if s.BorderBottom {
			h -= 1
		}
	}
	x += s.PaddingLeft
	y += s.PaddingTop
	w = int(w) - int(s.PaddingLeft) - int(s.PaddingRight)
	h = int(h) - int(s.PaddingTop) - int(s.PaddingBottom)
	if w < 0 {
		w = 0
	}
	if h < 0 {
		h = 0
	}
	return image.Rect(int(x), int(y), int(x)+w, int(y)+h)
}

func (s *widgetBlock) SetRect(x, y, w, h int) {
	s.X = x
	s.Y = y
	s.Width = w
	s.Height = h
}
