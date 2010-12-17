package main

import (
    "web"
    "kview"
    "os"
    "time"
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

var (
    menu = []MenuItem {
        MenuItem{"Home", "/"},
        MenuItem{"Edit", "/edit"},
    }
    global_ctx = struct{started string; hits uint} {
        time.LocalTime().Format("2006-01-02 15:04"),
        0,
    }
)

func home(web_ctx *web.Context) {
    req_ctx := MainCtx {
        title: "Home page",
        ctx:   Ctx {
            menu: Menu{menu, 0},
            left: []string {
                "This is a test service created entirely in Go (golang) " +
                "using <em>kasia.go</em>, <em>kview</em> and <em>web.go</em> " +
                "packages.",
                "Please select another menu item!",
            },
            right: RightCtx{"The House is much better than a flat. " +
                "So buy a new House today!"},
        },
    }
    global_ctx.hits++
    home_view.Exec(web_ctx, global_ctx, req_ctx)
}

func edit(web_ctx *web.Context) {
    req_ctx := MainCtx {
        title : "Edit page",
        ctx   : Ctx {
            menu:  Menu{menu, 1},
            left:  []string {
                "Hello! You can modify this example.",
                "Open <em>simple.go</em> file or some template file in your " +
                "editor and edit it.",
                "Then type: <code>$ make && ./simple</code>",
            },
            right: RightCtx{
                "To modify this example you may download " +
                "<a href='http://github.com/mikhailt/tabby'>tabby</a> source " +
                "editor writen entirely in Go!",
            },
        },
    }
    global_ctx.hits++
    edit_view.Exec(web_ctx, global_ctx, req_ctx)
}

// Init and run

func main() {
    if len(os.Args) != 1 {
        chrootuid()
    }

    // Change kview default template directory and error handler
    //kview.TemplatesDir = "some_dir"
    //kview.ErrorExit = new_error_handler

    viewInit()
    web.Get("/", home)
    web.Get("/edit", edit)
    web.Config.StaticDir = "static"
    web.Run("0.0.0.0:9999")
}
