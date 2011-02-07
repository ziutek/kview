package main

import "github.com/ziutek/kview"

var home_view, edit_view kview.View

func viewInit() {
    layout := kview.New("layout.kt")
    menu   := kview.New("menu.kt")

    // Create right column
    right  := kview.New("right.kt")
    // Add view components
    right.Div("Info",       kview.New("right/info.kt"))
    right.Div("Commercial", kview.New("right/commercial.kt"))

    // Create home view as layout copy.
    home_view = layout.Copy()
    home_view.Div("Menu",  menu)
    home_view.Div("Left",  kview.New("left/home.kt"))
    home_view.Div("Right", right)

    // Create edit view.
    edit_view = layout.Copy()
    edit_view.Div("Menu",  menu)
    edit_view.Div("Left",  kview.New("left/edit.kt"))
    edit_view.Div("Right", right)
}
