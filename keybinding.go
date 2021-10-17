// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

import "strings"

type keybinding struct {
	sequence string
	handler  func()
}

func (b *keybinding) match(ev Event) bool {
	return strings.EqualFold(b.sequence, ev.ID)
}
