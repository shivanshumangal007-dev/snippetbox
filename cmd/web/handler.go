package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"snippetbox.shivanshu.in/internal/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.servererror(w, err)
		return
	}
	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, http.StatusOK, "home.tmpl.html", data)
	// w.Write([]byte("hello from home route \n"))
}
func (app *application) Snippetview(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	if id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.servererror(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet
	// Create an instance of a templateData struct holding the snippet data.
	app.render(w, http.StatusOK, "view.tmpl.html", data)
}

// response. We'll update this shortly to show a HTML form.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.servererror(w, err)
		return
	}
	// Update the redirect path to use the new clean URL format.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
func (app *application) GetSnippets(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.servererror(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)

	}
}
