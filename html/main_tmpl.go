package main

import (
	"net/http"
	"text/template"
)

func main() {
	server := &http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("assets/tmpl.html")
	//t.Execute(w, "Hello World!")

	//t := template.New("assets/tmpl.html")
	//t, _ = t.ParseFiles("assets/tmpl.html")
	//t.Execute(w, "Hello World!")

	t, _ := template.ParseGlob("assets/*.html")
	t.Execute(w, "Hello World!")
}
