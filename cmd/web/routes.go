package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet/view", app.Snippetview)
	mux.HandleFunc("/snippet/create", app.Snippetcreate)
	mux.HandleFunc( "/snippets", app.GetSnippets)
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	return mux
}
