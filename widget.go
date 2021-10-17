// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

import (
	"github.com/cjbassi/gotop/colorschemes"
	"image"
	"sync"
)

// Alignment is used to set the direction in which widgets are laid out.
type LayoutMode int

// Available alignment options.
const (
	Horizontal LayoutMode = iota
	Vertical
)

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

// SizePolicy determines the space occupied by a widget.
type SizePolicy int

const (
	// Preferred interprets the size hint as the preferred size.
	Preferred SizePolicy = iota
	// Minimum allows the widget to shrink down to the size hint.
	Minimum
	// Maximum allows the widget to grow up to the size hint.
	Maximum
	// Expanding makes the widget expand to the available space.
	Expanding
)

type Widget interface {
	GetOuter() image.Rectangle
	GetInner() image.Rectangle
	GetInnerRealPos() image.Rectangle
	SetRect(x, y, w, h int)
	Draw()
	sync.Locker
	SetFocused(bool)
	IsFocused() bool
	SetActive(a bool)
	IsActive() bool
	SetParent(p Widget)
	GetPainter() *Painter
	SetPainter(p *Painter)
	DoEvent(e Event) bool
	OnEvent(fn func(w Widget, e *Event))
	Append(w Widget)
	Prepend(w Widget)
	Insert(i int, w Widget)
	Remove(i int)
	Length() int
	//SetBorder(enabled bool)
	//SetTitle(title string)
	SetText(text string)
	LayoutMode() LayoutMode               //子元素排列布局
	SizePolicy() (SizePolicy, SizePolicy) //显示模式
	MinSizeHint() image.Point
	SizeHint() image.Point
	Resize(pos image.Point, size image.Point)
	ReLayout() //重新布局
	sync.Locker
}

var colorscheme = colorschemes.Default

type Space struct {
	W uint
	H uint
}

type WidgetBase struct {
	widgetBlock
	active      bool //是否激活
	focused     bool
	children    []Widget
	layout      LayoutMode
	size        image.Point //不设置
	sizePolicyX SizePolicy
	sizePolicyY SizePolicy
	onEvent     func(w Widget, e *Event)
	text        string
	style       Style
	parent      Widget
	painter     *Painter
	px, py      int //用于记录当前的点输出点

	sync.Mutex
}

func (s *WidgetBase) Init() {

}

func (s *WidgetBase) GetParentMin() image.Point {
	if s.parent != nil {
		return s.parent.GetInnerRealPos().Min
	}
	return image.Pt(0, 0)
}

func (s *WidgetBase) draw() {
	s.px = 0
	s.py = 0
	if len(s.children) == 0 {
		s.drawClear()
	}
}

func (s *WidgetBase) Draw() {
	s.Lock()
	defer s.Unlock()
	s.draw()
	p := s.GetPainter()
	if p == nil {
		return
	}
	//drawSub
	s.drawSubWidget()
}

//刷新
func (s *WidgetBase) rePaint(w Widget) {
	p := s.GetPainter()
	if p == nil {
		return
	}
	p.addPaint(w)
	//p.Repaint(w)
}

func (s *WidgetBase) drawSubWidget() {
	for _, v := range s.children {
		v.Draw()
	}
}

func (s *WidgetBase) drawClear() {
	p := s.GetPainter()
	if p == nil {
		return
	}
	outer := s.GetOuter()
	min := s.parent.GetInnerRealPos().Min
	fillRect := outer.Add(min)
	p.Fill(Cell{' ', s.style}, fillRect)
}

//返回ture  消息将不在冒泡
func (s *WidgetBase) DoEvent(e Event) bool {
	switch e.Type {
	case KeyboardEvent:
		for _, child := range s.children {
			if child.DoEvent(e) {
				return true
			}
		}
	case MouseEvent:
		for _, child := range s.children {
			if child.DoEvent(e) {
				return true
			}
		}
	case ResizeEvent:
		payload := e.Payload.(Resize)
		termWidth, termHeight := payload.Width, payload.Height
		s.Resize(image.Point{0, 0}, image.Pt(termWidth, termHeight))
	}
	return false
}

func (s *WidgetBase) OnEvent(fn func(w Widget, e *Event)) {
	s.onEvent = fn
}

// SetFocused focuses the widget.
func (w *WidgetBase) SetFocused(f bool) {
	w.focused = f
}

// IsFocused returns whether the widget is focused.
func (w *WidgetBase) IsFocused() bool {
	return w.focused
}

// SetActive active the widget.
func (w *WidgetBase) SetActive(a bool) {
	w.active = a
}

// IsActive returns whether the widget is active.
func (w *WidgetBase) IsActive() bool {
	return w.active
}

// SetActive active the widget.
func (w *WidgetBase) SetParent(p Widget) {
	w.parent = p
}

// SetText
func (w *WidgetBase) SetText(text string) {
	w.text = text
}

// IsActive returns whether the widget is active.
func (w *WidgetBase) SetStyle(style Style) {
	w.style = style
}

// SetWidth returns whether the widget is active.
func (w *WidgetBase) SetWidth(width int) {
	w.size.X = width
	if w.size.X != 0 {
		w.sizePolicyX = Minimum
	} else {
		w.sizePolicyX = Expanding
	}
}

// SetHeight returns whether the widget is active.
func (w *WidgetBase) SetHeight(height int) {
	w.size.Y = height
	if w.size.Y != 0 {
		w.sizePolicyY = Minimum
	} else {
		w.sizePolicyY = Expanding
	}
}

func (w *WidgetBase) GetPainter() *Painter {
	if w.painter == nil {
		if w.parent != nil {
			w.painter = w.parent.GetPainter()
			return w.painter
		}
		return nil
	} else {
		return w.painter
	}
}

func (w *WidgetBase) SetPainter(p *Painter) {
	w.painter = p
}

// Append adds the given widget at the end of the Box.
func (s *WidgetBase) Append(w Widget) {
	s.children = append(s.children, w)
	w.SetParent(s)
}

// Prepend adds the given widget at the start of the Box.
func (s *WidgetBase) Prepend(w Widget) {
	s.children = append([]Widget{w}, s.children...)
	w.SetParent(s)
}

// Insert adds the widget into the Box at a given index.
func (s *WidgetBase) Insert(i int, w Widget) {
	if len(s.children) < i || i < 0 {
		return
	}

	s.children = append(s.children, nil)
	copy(s.children[i+1:], s.children[i:])
	s.children[i] = w
	w.SetParent(s)
}

// Remove deletes the widget from the Box at a given index.
func (s *WidgetBase) Remove(i int) {
	if len(s.children) <= i || i < 0 {
		return
	}

	s.children = append(s.children[:i], s.children[i+1:]...)
}

// Length returns the number of items in the box.
func (s *WidgetBase) Length() int {
	return len(s.children)
}

// Alignment returns the current alignment of the Box.
func (s *WidgetBase) LayoutMode() LayoutMode {
	return s.layout
}

// SizePolicy returns the current size policy.
func (w *WidgetBase) SizePolicy() (SizePolicy, SizePolicy) {
	return w.sizePolicyX, w.sizePolicyY
}

// MinSizeHint returns the minimum size hint for the layout.
func (s *WidgetBase) MinSizeHint() image.Point {
	var minSize image.Point
	if s.LayoutMode() == Horizontal {
		for _, child := range s.children {
			size := child.MinSizeHint()
			minSize.X += size.X
			if size.Y > minSize.Y {
				minSize.Y = size.Y
			}
		}
	} else {
		for _, child := range s.children {
			size := child.MinSizeHint()
			minSize.Y += size.Y
			if size.X > minSize.X {
				minSize.X = size.X
			}
		}
	}
	if s.size.X > 0 {
		minSize.X = s.size.X
	}
	if s.size.Y == 0 {
		minSize.Y = s.size.Y
	}
	return minSize
}

// SizeHint returns the recommended size hint for the layout.
func (s *WidgetBase) SizeHint() image.Point {
	var sizeHint image.Point

	for _, child := range s.children {
		size := child.SizeHint()
		if s.LayoutMode() == Horizontal {
			sizeHint.X += size.X
			if size.Y > sizeHint.Y {
				sizeHint.Y = size.Y
			}
		} else {
			sizeHint.Y += size.Y
			if size.X > sizeHint.X {
				sizeHint.X = size.X
			}
		}
	}

	return sizeHint
}

// Resize recursively updates the size of the Box and all the widgets it
// contains. This is a potentially expensive operation and should be invoked
// with restraint.
//
// Resize is called by the layout engine and is not intended to be used by end
// users.

func (s *WidgetBase) Resize(pos image.Point, size image.Point) {
	if size.X != s.Width || size.Y != s.Height || s.MarginLeft != pos.X || s.MarginTop != pos.Y {
		s.Width = size.X
		s.Height = size.Y
		s.X = pos.X
		s.Y = pos.Y
		inner := s.GetInner()
		s.layoutChildren(image.Pt(0, 0), inner.Size())
	}
}

//获取打印的真实坐标
func (s *WidgetBase) GetInnerRealPos() image.Rectangle {
	pMin := s.GetParentMin()
	inner := s.GetInner()
	return image.Rect(pMin.X+inner.Min.X, pMin.Y+inner.Min.Y, pMin.X+inner.Max.X, pMin.Y+inner.Max.Y)
}

//重新布局
func (s *WidgetBase) ReLayout() {
	inner := s.GetInner()
	s.layoutChildren(inner.Min, inner.Size())
}

func (s *WidgetBase) layoutChildren(pos image.Point, size image.Point) {
	space := doLayout(s.children, dim(s.LayoutMode(), size), s.LayoutMode())
	for i, sp := range space {
		switch s.LayoutMode() {
		case Horizontal:
			s.children[i].Resize(pos, image.Point{sp, size.Y})
			pos.X += sp
		case Vertical:
			s.children[i].Resize(pos, image.Point{size.X, sp})
			pos.Y += sp
		}
	}
}

func dim(a LayoutMode, pt image.Point) int {
	if a == Horizontal {
		return pt.X
	}
	return pt.Y
}

func alignedSizePolicy(a LayoutMode, w Widget) SizePolicy {
	hpol, vpol := w.SizePolicy()
	if a == Horizontal {
		return hpol
	}
	return vpol
}

func doLayout(ws []Widget, space int, a LayoutMode) []int {
	sizes := make([]int, len(ws))

	if len(sizes) == 0 {
		return sizes
	}

	remaining := space

	// Distribute MinSizeHint
	for {
		var changed bool
		for i, sz := range sizes {
			minSize := ws[i].MinSizeHint()
			if sz < dim(a, minSize) {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Distribute Minimum
	for {
		var changed bool
		for i, sz := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if p == Minimum && sz < dim(a, ws[i].SizeHint()) {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Distribute Preferred
	for {
		var changed bool
		for i, sz := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if (p == Preferred || p == Maximum) && sz < dim(a, ws[i].SizeHint()) {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Distribute Expanding
	for {
		var changed bool
		for i, sz := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if p == Expanding {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Distribute remaining space
	for {
		min := maxInt
		for i, s := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if (p == Preferred || p == Minimum) && s <= min {
				min = s
			}
		}
		var changed bool
		for i, sz := range sizes {
			if sz != min {
				continue
			}
			p := alignedSizePolicy(a, ws[i])
			if p == Preferred || p == Minimum {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

Resize:

	return sizes
}
