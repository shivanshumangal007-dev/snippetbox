package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	fmt.Println("hello world how are you man i am good")
	addr := flag.String("addr", ":4000", "addres of the the host")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile)

	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}
	app := &application{
		errorLog: errLog,
		infoLog:  infolog,
	}
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet/view", app.Snippetview)
	mux.HandleFunc("/snippet/create", app.Snippetcreate)
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	infolog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
