package kview

import (
    "os"
    "io"
    "log"
    "path"
    "github.com/ziutek/kasia.go"
)

var (
    // You can modify this, if you store templates in a different directory.
    TemplatesDir = "templates"

    // You can modify this, if want a different error handling.
    ErrorHandler = func(name string, err os.Error) {
        log.Printf("%%View '%s' error. %s\n", name, err.String())
    }
)

type View interface {
    Copy() View
    Strict(bool)
    Div(string, View)
    Exec(io.Writer, ...interface{})
    Render(...interface{}) *kasia.NestedTemplate
}

// View definition
type KView struct {
    name string
    tpl  *kasia.Template
    divs map[string]View
}

// Returns a pointer to a page
func New(name string) *KView {
    var (
        pg  KView
        err os.Error
    )
    pg.name = name
    pg.tpl, err = kasia.ParseFile(path.Join(TemplatesDir, name))
    if err != nil {
        ErrorHandler(name, err)
    }
    pg.divs = make(map[string]View)
    return &pg
}

// Returns a pointer to a copy of the page
func (pg *KView) Copy() View {
    new_pg := *pg
    // Make a copy of divs map
    new_pg.divs = make(map[string]View)
    for k, v := range pg.divs {
        new_pg.divs[k] = v
    }
    return &new_pg
}

// Set strig render flag
func (pg *KView) Strict(strict bool) {
    pg.tpl.Strict = strict
}

// Add subview
func (pg *KView) Div(name string, view View) {
    pg.divs[name] = view
}

func prepend(slice []interface{}, pre ...interface{}) (ret []interface{}) {
    ret = make([]interface{}, len(slice) + len(pre))
    copy(ret, pre)
    copy(ret[len(pre):], slice)
    return
}

// Render view to wr with data
func (pg *KView) Exec(wr io.Writer, ctx ...interface{}) {
    // Add divs to the bottom of the context stack
    ctx = prepend(ctx, pg.divs)
    err := pg.tpl.Run(wr, ctx...)
    if err != nil {
        ErrorHandler(pg.name, err)
    }
}

// Use this method in template text to render page inside other page.
func (pg *KView) Render(ctx ...interface{}) *kasia.NestedTemplate {
    if len(ctx) > 0 {
        // Check if render was called with full template context as first arg
        if ci, ok := ctx[0].(kasia.ContextItself); ok {
            // Rearange context, remove old divs
            ctx = append(ci[1:], ctx[1:])
        }
    }
    // Add divs to the bottom of the context stack
    ctx = prepend(ctx, pg.divs)
    return pg.tpl.Nested(ctx...)
}
