// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

type Box struct {
	Block
}

func NewHBox(c ...Widget) *Box {
	b := &Box{
		Block: *NewBlock(),
	}
	for _, v := range c {
		b.Append(v)
	}
	b.layout = Horizontal
	b.style = NewStyle(ColorWhite)
	return b
}

func NewVBox(c ...Widget) *Box {
	b := &Box{
		Block: *NewBlock(),
	}
	for _, v := range c {
		b.Append(v)
	}
	b.layout = Vertical
	b.style = NewStyle(ColorWhite)
	return b
}
