package main

import "github.com/ziutek/kview"

var home_view, edit_view kview.View

func viewInit() {
    layout := kview.New("layout.kt")
    menu   := kview.New("menu.kt")

    // Create right column
    right  := kview.New("right.kt")
    // Add view components
    right.Div("info",       kview.New("right/info.kt"))
    right.Div("commercial", kview.New("right/commercial.kt"))

    // Create home view as layout copy.
    home_view = layout.Copy()
    home_view.Div("menu",  menu)
    home_view.Div("left",  kview.New("left/home.kt"))
    home_view.Div("right", right)

    // Create edit view.
    edit_view = layout.Copy()
    edit_view.Div("menu",  menu)
    edit_view.Div("left",  kview.New("left/edit.kt"))
    edit_view.Div("right", right)
}
