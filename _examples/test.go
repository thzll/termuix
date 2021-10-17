// Copyright 2021. thzll <tanghuizll@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	uix "github.com/thzll/termuix"
	"log"
)

func main() {
	l1 := uix.NewLabel("我xxxxxxxxxxxxxxxxxxxxxx")
	//l1.SetWidth(10)
	l2 := uix.NewLabel("yyyyyyyyyyyyyyyyyyyyyy")
	input := uix.NewInput()
	input.SetHeight(1)
	input.SetText("hello")
	input.OnSubmit(func(input *uix.Input) {
		l2.SetText(input.Text())
		input.SetText("")
	})
	input.SetHeight(3)
	input.Border = true
	input.BorderTop = true
	input.BorderBottom = true
	h1 := uix.NewHBox()
	//h1.BorderLeft = true
	//h1.BorderTop = true
	//h1.BorderBottom = true
	//h1.WidgetBase.Max = image.Pt(100,100)
	//h1.SetBorder(true)
	h1.SetTitle("∞∞∞∞∞∞ chat ∞∞∞∞∞∞∞")
	v1 := uix.NewVBox()
	v2 := uix.NewVBox()
	//v1.SetBorder(false)
	v1.Append(l1)
	v2.Append(l2)
	v2.Append(input)
	h1.Append(v1)
	h1.Append(v2)
	ui, err := uix.New(h1)
	if err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Quit()
	ui.Run()
}
