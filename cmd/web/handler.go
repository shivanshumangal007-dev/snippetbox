package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"snippetbox.shivanshu.in/internal/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Print(err.Error())
		app.servererror(w, err)
		return
	}
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.servererror(w, err)
		return
	}
	data := &templateData{
		Snippets: snippets,
	}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.errorLog.Print(err.Error())
		app.servererror(w, err)
		return
	}
	// w.Write([]byte("hello from home route \n"))
}
func (app *application) Snippetview(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.servererror(w, err)
		return
	}
	// Create an instance of a templateData struct holding the snippet data.
	data := &templateData{
		Snippet: snippet,
	}
	// Pass in the templateData struct when executing the template.
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.servererror(w, err)
	}
}
func (app *application) Snippetcreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("cerate your new snippet here \n"))
	id, err := app.snippets.Insert("first snippet", "lorem ipsum hell you go man i lov eyou hello man", 7)
	if err != nil {
		app.servererror(w, err)
	}
	app.infoLog.Printf("id of the created snippets is %d", id)
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
