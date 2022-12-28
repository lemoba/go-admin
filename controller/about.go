package controller

import (
	"net/http"
	"text/template"
)

func registerAboutRoutes() {
	http.HandleFunc("/about", handleAbout)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/layout.html", "assets/about.html")
	t.ExecuteTemplate(w, "layout", "about")
}
