package controller

import (
	"net/http"
	"text/template"
)

func registerHomeRoutes() {
	http.HandleFunc("/home", handleHome)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/layout.html", "assets/home.html")
	t.ExecuteTemplate(w, "layout", "home")
}
