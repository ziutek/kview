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

// View definition
type View struct {
    name string
    tpl  *kasia.Template
    Divs map[string]*View
}

// Returns a pointer to a page
func New(name string) *View {
    var (
        pg  View
        err os.Error
    )
    pg.name = name
    pg.tpl, err = kasia.ParseFile(path.Join(TemplatesDir, name))
    if err != nil {
        ErrorExit(name, err)
    }
    pg.Divs = make(map[string]*View)
    return &pg
}

// Returns a pointer to a copy of the page
func (pg *View) Copy() *View {
    new_pg := *pg
    // Make a copy of Divs map
    new_pg.Divs = make(map[string]*View)
    for k, v := range pg.Divs {
        new_pg.Divs[k] = v
    }
    return &new_pg
}

// Set strig render flag
func (pg *View) Strict(strict bool) {
    pg.tpl.Strict = strict
}

func prepend(slice []interface{}, pre ...interface{}) (ret []interface{}) {
    ret = make([]interface{}, len(slice) + len(pre))
    copy(ret, pre)
    copy(ret[len(pre):], slice)
    return
}

// Render view to wr with data
func (pg *View) Exec(wr io.Writer, ctx ...interface{}) {
    ctx = prepend(ctx, pg.Divs)
    err := pg.tpl.Run(wr, ctx...)
    if err != nil {
        ErrorExit(pg.name, err)
    }
}

// Use this method in template text to render page inside other page.
func (pg *View) Render(ctx ...interface{}) *kasia.NestedTemplate {
    ctx = prepend(ctx, pg.Divs)
    return pg.tpl.Nested(ctx...)
}
