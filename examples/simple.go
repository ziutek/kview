package main

import (
    "os"
    "time"
    "fmt"
    "github.com/hoisie/web.go"
    "github.com/ziutek/kview"
)

type MenuItem struct {
    Name, Url string
}

type Menu struct {
    Content  []MenuItem
    Selected int
}

type RightCtx struct {
    Commercial string
}

type Ctx struct {
    // Title of the page
    Title string
    // Navigation menu
    Menu  Menu
    // Data for the left column of the page
    Left  interface{}
    // Data for the right column of the page
    Right RightCtx
}

var (
    menu = []MenuItem {
        MenuItem{"Home", "/"},
        MenuItem{"Edit", "/edit"},
    }
    // Some global variables presented on the Web
    global_ctx = struct{Started, LastCliAddr string; Hits uint} {
        time.LocalTime().Format("2006-01-02 15:04"),
        "",
        0,
    }
)

// Renders view and actualizes global context
func exec(web_ctx *web.Context, view kview.View, req_ctx interface{}) {
    global_ctx.Hits++
    view.Exec(web_ctx, global_ctx, req_ctx)
    global_ctx.LastCliAddr = web_ctx.RemoteAddr
}

// The home page handler
func home(web_ctx *web.Context) {
    req_ctx := Ctx {
        Title: "Home page",
        Menu: Menu{menu, 0},
        Left: []string {
            "This is a test service created entirely in Go (golang) using " +
            "<em>kasia.go</em>, <em>kview</em> and <em>web.go</em> packages.",
            "Please select another menu item!",
        },
        Right: RightCtx{"A house is much better than a flat. So buy a new " +
            "House today!"},
    }
    exec(web_ctx, home_view, req_ctx)
}

// The Edit page handler
func edit(web_ctx *web.Context) {
    req_ctx := Ctx {
        Title : "Edit page",
        Menu:  Menu{menu, 1},
        Left:  []string {
            "Hello! You can modify this example.",
            "Open <em>simple.go</em> file or some template file in your " +
            "editor and edit it.",
            "Then type: <code>$ make && ./simple</code>",
        },
        Right: RightCtx{
            "To modify this example you may download " +
            "<a href='http://github.com/mikhailt/tabby'>tabby</a> source " +
            "editor writen entirely in Go!",
        },
    }
    exec(web_ctx, edit_view, req_ctx)
}

// Init and run

func main() {
    if len(os.Args) == 3 {
        chrootuid(os.Args[1], os.Args[2])
    } else if len(os.Args) != 1 {
        fmt.Printf("Usage: %s [DIRECTORY UID]\n", os.Args[0])
        os.Exit(1)
    }

    // Change kview default template directory and error handler
    //kview.TemplatesDir = "some_dir"
    //kview.ErrorHandler = new_error_handler

    viewInit()
    web.Get("/", home)
    web.Get("/edit", edit)
    web.Config.StaticDir = "static"
    web.Run("0.0.0.0:9999")
}
