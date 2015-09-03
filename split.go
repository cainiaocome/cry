// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

var theme gxui.Theme
func config_field(l string, d string) gxui.LinearLayout {  // l:label d:default
    linear_layout := theme.CreateLinearLayout()
    linear_layout.SetDirection(gxui.LeftToRight)
    label := theme.CreateLabel()
    label.SetText(l + ": ")
    textbox :=  theme.CreateTextBox()
    textbox.SetText(d)
    linear_layout.AddChild(label)
    linear_layout.AddChild(textbox)
    return linear_layout
}
func apanel() gxui.PanelHolder{
	holder := theme.CreatePanelHolder()
    config_linear_layout := theme.CreateLinearLayout()
    config_linear_layout.SetDirection(gxui.TopToBottom)

    //thread
    thread_linear_layout := config_field("thread", "1000")
    //ip
    ip_linear_layout := config_field("ip", "ip.txt")
    //username
    username_linear_layout := config_field("username", "username.txt")
    //password
    password_linear_layout := config_field("password", "password.txt")
    //good result
    good_linear_layout := config_field("output", "good.txt")
    //loglevel
    loglevel_linear_layout := config_field("loglevel", "1")

    config_linear_layout.AddChild(thread_linear_layout)
    config_linear_layout.AddChild(ip_linear_layout)
    config_linear_layout.AddChild(username_linear_layout)
    config_linear_layout.AddChild(password_linear_layout)
    config_linear_layout.AddChild(good_linear_layout)
    config_linear_layout.AddChild(loglevel_linear_layout)
	holder.AddPanel(config_linear_layout, "config")
	return holder
}
func bpanel() gxui.PanelHolder {
	holder := theme.CreatePanelHolder()
    label := theme.CreateLabel()
    label.SetText("Wrote By Alex.\n\nUse It At Your Own Rist")
	holder.AddPanel(label, "disclosure")
	return holder
}
func cpanel() gxui.PanelHolder {
	holder := theme.CreatePanelHolder()
    label := theme.CreateLabel()
    label.SetText("")
	holder.AddPanel(label, "log")
	return holder
}
func appMain(driver gxui.Driver) {
    theme = dark.CreateTheme(driver)

	// ┌───────┐║┌───────┐
	// │       │║│       │
	// │   A   │║│   B   │
	// │       │║│       │
	// └───────┘║└───────┘
	// ═══════════════════
	// ┌─────────────────┐
	// │                 │
	// │   C             │
	// │                 │
	// └─────────────────┘

	splitterAB := theme.CreateSplitterLayout()
	splitterAB.SetOrientation(gxui.Horizontal)
	splitterAB.AddChild(apanel())
	splitterAB.AddChild(bpanel())

	splitterC := theme.CreateSplitterLayout()
	splitterC.SetOrientation(gxui.Horizontal)
	splitterC.AddChild(cpanel())

	vSplitter := theme.CreateSplitterLayout()
	vSplitter.SetOrientation(gxui.Vertical)
	vSplitter.AddChild(splitterAB)
	vSplitter.AddChild(splitterC)

	window := theme.CreateWindow(800, 600, "CRY")
	window.AddChild(vSplitter)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
