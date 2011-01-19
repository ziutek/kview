package kview

import (
    "os"
    "io"
    "log"
    "path"
    "kasia"
)

var (
    // You can modify this, if you store templates in a different directory.
    TemplatesDir = "templates"

    // You can modify this, if want a different error handling.
    ErrorExit = func(name string, err os.Error) {
        log.Exitf("%%View '%s' error. %s\n", name, err.String())
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
        ErrorExit(name, err)
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
    //ctx = prepend(ctx, pg.divs)
    ctx = append(ctx, pg.divs)
    err := pg.tpl.Run(wr, ctx...)
    if err != nil {
        ErrorExit(pg.name, err)
    }
}

// Use this method in template text to render page inside other page.
func (pg *KView) Render(ctx ...interface{}) *kasia.NestedTemplate {
    if len(ctx) > 0 {
        if ci, ok := ctx[0].(kasia.ContextItself); ok {
            // Render was called with full template context as first argument
            ci[0] = pg.divs // Change divs
            // Append other parameters to the context before call Nested
            return pg.tpl.Nested(append(ci, ctx[1:]...)...)
        }
    }
    //ctx = prepend(ctx, pg.divs)
    ctx = append(ctx, pg.divs)
    return pg.tpl.Nested(ctx...)
}
