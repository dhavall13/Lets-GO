package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Handler

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request)  {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request)  {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST") // first parameter is header name and second is the value
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed )

		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main()  {
	// servermux = router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the http.ListenAndServe() function to start a new web server
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}