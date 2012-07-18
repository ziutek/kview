## What is *kview*

Kview is simple but useful wrapper for
[Kasia.go](https://github.com/ziutek/kasia.go) templates, which helps to
modularize content of dynamic website. It allows you to easily describe the
relationship between modules of your website.

#### Build the structure of your web service

You can build a web page from blocks. Every block is associated with different
template file.

Example:

    // Web pages declared as global variables
    var home_view, edit_view kview.View

    // Create the web service hierarchy
    func webInit() {
        // Load site layout. 
        layout := kview.New("layout.kt")

        // Example layout consists of header, menu, two columns and footer. The
        // left column can contain very different information. Therefore, for
        // any page it will be defined by the different template. The right
        // column shows always the same type of information, so will be defined
        // once for all pages (the same applies to the menu).

        // Load the menu
        menu := kview.New("menu.kt")

        // Load the right column
        right  := kview.New("right.kt")

        // Add components to the right column
        right.Div("Info",       kview.New("right/info.kt"))
        right.Div("Commercial", kview.New("right/commercial.kt"))

        // Create the first page as layout copy. It is efficient operation
        // (references are copied, not the data itself).
        home_view = layout.Copy()

        // Add page components
        home_view.Div("Menu",  menu)
        home_view.Div("Left",  kview.New("left/home.kt", utils))
        home_view.Div("Right", right)

        // Create the second page.
        edit_view = layout.Copy()
        edit_view.Div("Menu",  menu)
        edit_view.Div("Left",  kview.New("left/edit.kt")
        edit_view.Div("Right", right)
    }

The structure of the service is ready. The (optional) *utils* variable used in
*Left* div may contains your utility functions/variables:

    var utils = map[string]interface{} {
        "contains": strings.Contains,
        "addf": func(a, b float64) float64 {return a + b},
		"pi": 3.14159,
    }

You can use them in *left/home.kt* template as follows:

    $a + $pi = $addf(a, pi)

    $if contains(s, "abc"):
        The s variable contains 'abc' substring.
    $else:
        The s doesn't contain 'abc' substring.
    $end

Some useful functions are provided by default:

* `len(interface{}) int` - it returns length of array/slice or -1,
* `fmt(format string, a ...interface{}) string` - works like *fmt.Sprintf*
  (in fact it is *fmt.Sprintf*),
* `add, sub, mul, div, mod (a, b int) int` - arithmetic operations,
* `and, or (a, b bool)`, `not(a bool) bool` - logical operations.


#### Publishing your web service

To publish the service created with *kview* you can use the *web.go* framework:

    func home(web_ctx *web.Context) {
        // ...
        // tpl_ctx contains the data needed to display the page
        home_view.Exec(web_ctx, tpl_ctx)
    }

    func edit(web_ctx *web.Context) {
        // ...
        edit_view.Exec(web_ctx, tpl_ctx)
    }

    func main() {
        webInit()

        web.Get("/", home)
        web.Get("/edit", edit)
        web.Run("0.0.0.0:80")
    }

the *http* package:

    func home(con http.ResponseWriter, req *http.Request) {
        // ...
        home_view.Exec(con, tpl_ctx) 
    }

    func edit(con http.ResponseWriter, req *http.Request) {
        // ...
        edit_view.Exec(con, tpl_ctx)
    }

    func main() {
        webInit()

        http.HandleFunc("/", home)
        http.HandleFunc("/edit", edit)
        // ...
        http.ListenAndServe("0.0.0.0:80", nil)
    }

or other available framework like *twister*.

In the template, nested blocks are visible under the name you given them. You
can render them using *Render* method (rather than *Nested* method in pure
*kasia.go*):

    <div id='Left'>$Left.Render(ctx.left)</div>
    <div id='Right'>$Right.Render(ctx.right)</div>

You can find a working example in the *examples* directory. This simple application consists of:

* three files with Go code (the *chrootuid.go* file is irrelevant for our
  considerations - contains code for change application root directory and its
  privileges),
* templates tree (in *templates* directory),
* CSS style sheet (in *static* directory).

## How to install and run example application

Instal [web.go](http://github.com/hoisie/web.go):

    $ goinstall github.com/hoisie/web.go

Next install *kview*:

    $ goinstall github.com/ziutek/kview

This command implicitly install [kasia.go](http://github.com/ziutek/kasia.go)
too.

Next you can build and run an example application:

    $ cd $GOROOT/src/pkg/github.com/ziutek/kview/examples
    $ make
    $ ./simple

Next launch your browser and open the URL: http://127.0.0.1:9999

You can try [this link](http://195.74.48.3:9999/) first, if you want to see a
sample application in action.
