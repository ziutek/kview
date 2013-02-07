// Wrapper for kasia.go templates designed for easy writing web applications.
package kview

import (
	"fmt"
	"github.com/ziutek/kasia.go"
	"io"
	"log"
	"path"
	"reflect"
	"strings"
)

var (
	// You can modify this, if you store templates in a different directory.
	TemplatesDir = "templates"

	// You can modify this, if you want a different error handling.
	ErrorHandler = func(name string, err error) {
		log.Printf("%%View '%s' error. %s\n", name, err.Error())
	}
)

// Globals can store useful things that are global for all views.
// You can add there your functions/variables which will be visible in any
// view. See also symbols parameter in New function.
var Globals = map[string]interface{}{
	"len": func(a interface{}) int {
		v := reflect.ValueOf(a)
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice || v.Kind() == reflect.String {
			return v.Len()
		}
		return -1
	},
	"fmt":  fmt.Sprintf,
	"join": strings.Join,
	"add": func(a, b int) int {
		return a + b
	},
	"sub": func(a, b int) int {
		return a - b
	},
	"mul": func(a, b int) int {
		return a * b
	},
	"div": func(a, b int) int {
		return a / b
	},
	"mod": func(a, b int) int {
		return a % b
	},
	"and": func(a, b bool) bool {
		return a && b
	},
	"or": func(a, b bool) bool {
		return a || b
	},
	"not": func(a bool) bool {
		return !a
	},
}

type View interface {
	Copy() View
	Strict(bool)
	Div(string, View)
	Exec(io.Writer, ...interface{})
	Render(...interface{}) *kasia.NestedTemplate
}

// View definition
type KView struct {
	name    string
	tpl     *kasia.Template
	symbols map[string]interface{}
}

// Returns a pointer to the view
func New(name string, symbols ...map[string]interface{}) *KView {
	var (
		pg  KView
		err error
	)
	pg.name = name
	pg.tpl, err = kasia.ParseFile(path.Join(TemplatesDir, name))
	if err != nil {
		ErrorHandler(name, err)
	}
	pg.symbols = make(map[string]interface{})
	for _, g := range symbols {
		for k, v := range g {
			pg.symbols[k] = v
		}
	}
	return &pg
}

// Returns a copy of the view
func (pg *KView) Copy() View {
	newpg := *pg
	// Make a copy of globals
	newpg.symbols = make(map[string]interface{})
	for k, v := range pg.symbols {
		newpg.symbols[k] = v
	}
	return &newpg
}

// Set the strict render flag
func (pg *KView) Strict(strict bool) {
	pg.tpl.Strict = strict
}

// Add subview
func (pg *KView) Div(name string, view View) {
	pg.symbols[name] = view
}

func prepend(slice []interface{}, pre ...interface{}) (ret []interface{}) {
	ret = make([]interface{}, 0, len(slice)+len(pre))
	ret = append(ret, pre...)
	ret = append(ret, slice...)
	return
}

// Render view to w with data
func (pg *KView) Exec(w io.Writer, ctx ...interface{}) {
	// Add globals to the bottom of the context stack
	ctx = prepend(ctx, Globals, pg.symbols)
	err := pg.tpl.Run(w, ctx...)
	if err != nil {
		ErrorHandler(pg.name, err)
	}
}

// Use this method in template text to render view inside other view.
func (pg *KView) Render(ctx ...interface{}) *kasia.NestedTemplate {
	if len(ctx) > 0 {
		// Check if render was called with full template context as first arg
		if ci, ok := ctx[0].(kasia.ContextItself); ok {
			// Rearange context
			newctx := make([]interface{}, 0, len(ci)+len(ctx)-1)
			newctx = append(newctx, ci...)
			newctx = append(newctx, ctx[1:]...)
			// Replace symbols
			newctx[1] = pg.symbols
			return pg.tpl.Nested(newctx...)
		}
	}
	// Add globals to the bottom of the context stack
	ctx = prepend(ctx, Globals, pg.symbols)
	return pg.tpl.Nested(ctx...)
}
