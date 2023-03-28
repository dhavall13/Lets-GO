package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)



func main()  {

	addr := flag.String("addr", ":4000", "HTTP network address")


	flag.Parse()

	// log.New to create logger for writing messages in cmd
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	// servermux = router
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the http.ListenAndServe() function to start a new web server
	infoLog.Printf("Starting server on %s", *addr) // information message
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err) // err message
}