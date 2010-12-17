## What is *kview*

Kview is simple wrapper for Kasia.go templates, which helps to modularize
content of dynamic website. It allows you to easily describe the relationship
between modules of your website.

You can build the web page from the blocks directly in your Go code. Every block
can be a template defined in separate file.

Example:

    // Load site layout. 

    layout := kview.New("layout.kt")
    
    // Layout consists of two columns. The left column can contain very
    // different information. Therefore, for any page it will be defined by
    // the different template. The right column shows always the same type of
    // information, so will be defined once for all pages. The same applies to
    // the menu.

    // Load the menu

    menu := kview.New("menu.kt")

    // Load the right column

    right  := kview.New("right.kt")

    // Add components of the right column

    right.Divs["Info"]       = kview.New("right/info.kt")
    right.Divs["Commercial"] = kview.New("right/commercial.kt")

    // Create the first page as layout copy. It is efficient operation
    // (references are copied, not the data itself).

    home_view = layout.Copy()

    // Add page components

    home_view.Divs["Menu"]  = menu
    home_view.Divs["Left"]  = kview.New("left/home.kt")
    home_view.Divs["Right"] = right

    // Create the second page.

    edit_view = layout.Copy()
    edit_view.Divs["Menu"]  = menu
    edit_view.Divs["Left"]  = kview.New("left/edit.kt")
    edit_view.Divs["Right"] = right    

The structure of the service is ready. You can publish it with web.go:

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
        // ...
        web.Get("/", home)
        web.Get("/edit", edit)
        web.Run("0.0.0.0:80")
    }
    
or http package:

    func home(con http.ResponseWriter, req *http.Request) {
        // ...
        home_view.Exec(con, tpl_ctx) 
    }

    func edit(con http.ResponseWriter, req *http.Request) {
        // ...
        edit_view.Exec(con, tpl_ctx)
    }

    func main() {
        // ...
        http.HandleFunc("/", home)
        http.HandleFunc("/edit", edit)
        // ...
        http.ListenAndServe("0.0.0.0:80", nil)
    }

In the template, nested blocks ar visable under the name you given them. You can
render them using *Render* method (rather than *Nested* method in pure
*kasia.go*):

    <div id='Left'>$Left.Render(ctx.left)</div>
    <div id='Right'>$Right.Render(ctx.right)</div>


You can find a working example (one file with Go code, template tree and CSS
style sheet) in the *examples* directory.

## How to install and run example application

Instal [web.go](http://github.com/hoisie/web.go) and
[kasia.go](http://github.com/ziutek/kasia.go) first. Next instal *kview* and
compile and run simple example application:

    $ git clone git://github.com/ziutek/kview
    $ cd kview && make install
    $ cd examples && make
    $ ./simple

Next launch your browser and open the URL: http://127.0.0.1:9999

You can try [this link](http://195.74.48.3:9999/) first, if you want to see a
sample application in action.
