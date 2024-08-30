package routes

import (
	"groupie-tracker-filters/loaders"
	"html/template"
	"net/http"
	"strings"
)

// index is a function to define the index route
func index() {
	http.HandleFunc("/", indexHandler)
}

// indexHandler is the handler for the index route
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Redirection.
	idPath := "/artist/"
	if r.URL.Path != "/" && !strings.HasPrefix(r.URL.Path, idPath) && r.URL.Path != idPath {
		http.Redirect(w, r, "/error404", http.StatusTemporaryRedirect)
		return
	}
	// template.
	pageTemplate, err := template.ParseFiles("client/index.html")
	if err != nil {
		http.Redirect(w, r, "/error500", http.StatusTemporaryRedirect)
		return
	}
	// Récupération des membres cochés
	members := r.URL.Query()["members"]
	
	// Appel de loaders.Index avec les arguments corrects
	loaders.Index(w, r, pageTemplate, members)
}
