// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package component

import "github.com/gizak/termui/v3"

type Box struct {
	WidgetBase
}

func NewHBox() *Box {
	return &Box{
		WidgetBase{
			layout:      Horizontal,
			Block:       *termui.NewBlock(),
			sizePolicyX: Expanding,
			sizePolicyY: Expanding,
		},
	}
}

func NewVBox() *Box {
	return &Box{
		WidgetBase{
			layout:      Vertical,
			Block:       *termui.NewBlock(),
			sizePolicyX: Expanding,
			sizePolicyY: Expanding,
		},
	}
}
