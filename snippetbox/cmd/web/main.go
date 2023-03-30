package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}

func main()  {

	addr := flag.String("addr", ":4000", "HTTP network address")


	flag.Parse()

	// log.New to create logger for writing messages in cmd
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	app := application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	// servermux = router
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
        Addr:     *addr,
        ErrorLog: errorLog,
        Handler:  mux,
    }

	// Use the http.ListenAndServe() function to start a new web server
	infoLog.Printf("Starting server on %s", *addr) // information message
	err := srv.ListenAndServe()
	errorLog.Fatal(err) // err message
}