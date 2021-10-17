// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"github.com/cjbassi/gotop/colorschemes"
	_ "github.com/cjbassi/gotop/colorschemes"
	ui "github.com/gizak/termui/v3"
)

var colorscheme = colorschemes.Default

func SetDefaultTermuiColors() {
	ui.Theme.Default = ui.NewStyle(ui.Color(colorscheme.Fg), ui.Color(colorscheme.Bg))
	ui.Theme.Block.Title = ui.NewStyle(ui.Color(colorscheme.BorderLabel), ui.Color(colorscheme.Bg))
	ui.Theme.Block.Border = ui.NewStyle(ui.Color(colorscheme.BorderLine), ui.Color(colorscheme.Bg))
}
