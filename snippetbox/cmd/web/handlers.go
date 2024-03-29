package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

	files := []string{
        "./ui/html/base.html",
        "./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
    }


	ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

	err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        app.serverError(w, err)
    }

}

// Add a snippetView handler function.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request)  {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request)  {

	if r.Method !=http.MethodPost {
		w.Header().Set("Allow", http.MethodPost) // first parameter is header name and second is the value
		app.clientError(w, http.StatusMethodNotAllowed)

		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}