package main

import (
    "web"
    "kview"
)


var home_view, edit_view *kview.View

func viewInit() {
    layout := kview.New("layout.kt")
    menu   := kview.New("menu.kt")

    // Create right column
    right  := kview.New("right.kt")
    // Add view components
    right.Divs["Info"]       = kview.New("right/info.kt")
    right.Divs["Commercial"] = kview.New("right/commercial.kt")

    // Create home view as layout copy.
    home_view = layout.Copy()
    home_view.Divs["Menu"]  = menu
    home_view.Divs["Left"]  = kview.New("left/home.kt")
    home_view.Divs["Right"] = right

    // Create edit view.
    edit_view = layout.Copy()
    edit_view.Divs["Menu"] = menu
    edit_view.Divs["Left"] = kview.New("left/edit.kt")
    edit_view.Divs["Right"] = right
}


type MenuItem struct {
    name, url string
}

type Menu struct {
    content  []MenuItem
    selected int
}

type RightCtx struct {
    commercial string
}

type Ctx struct {
    menu  Menu
    left  interface{}
    right RightCtx
}

type MainCtx struct {
    title string
    ctx   Ctx
}

var menu = []MenuItem{
    MenuItem{"Home", "/"},
    MenuItem{"Edit", "/edit"},
}

func home(web_ctx *web.Context) {
    tpl_ctx := MainCtx {
        title: "Home page",
        ctx:   Ctx {
            menu: Menu{menu, 0},
            left: []string {
                "This is a test service created with <em>kasia.go</em> " +
                "and <em>kview</em>.",
                "Please select another menu item!",
            },
            right: RightCtx{"Buy your new Home!"},
        },
    }
    home_view.Exec(web_ctx, tpl_ctx)
}

func edit(web_ctx *web.Context) {
    tpl_ctx := MainCtx {
        title : "Edit page",
        ctx   : Ctx {
            menu:  Menu{menu, 1},
            left:  "Hello! You can modify this example." +
                "This text is in <em>simple.go</em> file.",
            right: RightCtx{"Buy new great Goedit editor!"},
        },
    }
    edit_view.Exec(web_ctx, tpl_ctx)
}

// Init and run

func main() {
    // Change kview default template directory and error handler
    //kview.TemplatesDir = "some_dir"
    //kview.ErrorExit = new_error_handler

    viewInit()
    web.Get("/", home)
    web.Get("/edit", edit)
    web.Run("0.0.0.0:9999")
}
