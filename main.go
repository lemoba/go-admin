package main

import (
	"html/template"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/tmpl.html")
	t.Execute(w, "Hello World!")
}
