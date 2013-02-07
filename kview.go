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
	globals map[string]interface{}
}

// Returns a pointer to the page
func New(name string, globals ...map[string]interface{}) *KView {
	var (
		pg  KView
		err error
	)
	pg.name = name
	pg.tpl, err = kasia.ParseFile(path.Join(TemplatesDir, name))
	if err != nil {
		ErrorHandler(name, err)
	}
	pg.globals = make(map[string]interface{})
	// First some default globals
	for k, v := range Globals {
		pg.globals[k] = v
	}
	// globals may default
	for _, g := range globals {
		for k, v := range g {
			pg.globals[k] = v
		}
	}
	return &pg
}

// Returns a copy of the page
func (pg *KView) Copy() View {
	new_pg := *pg
	// Make a copy of globals
	new_pg.globals = make(map[string]interface{})
	for k, v := range pg.globals {
		new_pg.globals[k] = v
	}
	return &new_pg
}

// Set the strict render flag
func (pg *KView) Strict(strict bool) {
	pg.tpl.Strict = strict
}

// Add subview
func (pg *KView) Div(name string, view View) {
	pg.globals[name] = view
}

func prepend(slice []interface{}, pre ...interface{}) (ret []interface{}) {
	ret = make([]interface{}, len(slice)+len(pre))
	copy(ret, pre)
	copy(ret[len(pre):], slice)
	return
}

// Render view to wr with data
func (pg *KView) Exec(wr io.Writer, ctx ...interface{}) {
	// Add globals to the bottom of the context stack
	ctx = prepend(ctx, pg.globals)
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
			// Rearange context, remove old globals
			ctx = append(ci[1:], ctx[1:])
		}
	}
	// Add globals to the bottom of the context stack
	ctx = prepend(ctx, pg.globals)
	return pg.tpl.Nested(ctx...)
}

// Some useful functions for globals.
// You can add there your functions/variables which will be visable in any
// view. See also globals parameter in New function.
var Globals = map[string]interface{}{
	"len": func(a interface{}) int {
		v := reflect.ValueOf(a)
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice || v.Kind() == reflect.String {
			return v.Len()
		}
		return -1
	},
	"fmt": fmt.Sprintf,
	"join" : strings.Join,
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
