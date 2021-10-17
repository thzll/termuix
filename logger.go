// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termuix

import (
	"io/ioutil"
	"log"
)

var logger Logger = log.New(ioutil.Discard, "", 0)

// Logger provides a interface for the standard logger.
type Logger interface {
	Printf(format string, args ...interface{})
}

// SetLogger sets the logger that is used in tui.
func SetLogger(l Logger) {
	logger = l
}
