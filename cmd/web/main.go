package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hello world how are you man i am good")
	addr := flag.String("addr", ":4000", "addres of the the host")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static"))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet/view", Snippetview)
	mux.HandleFunc("/snippet/create", Snippetcreate)
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	infolog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
