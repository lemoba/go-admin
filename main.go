package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

type aboutHandler struct{}

func (a *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
func main() {
	//mh := &helloHandler{}
	//ah := &aboutHandler{}
	//
	//server := http.Server{
	//	Addr:    "localhost:8080",
	//	Handler: nil, // DefaultServeMux
	//}
	//http.Handle("/hello", mh)
	//http.Handle("/about", ah)
	//
	//http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("home"))
	//})

	// http.HandleFunc("/welcome", welcome)
	// http.Handle("/welcome", http.HandlerFunc(welcome))

	// server.ListenAndServe()
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "assets"+r.URL.Path)
	//})

	// http.ListenAndServe("localhost:8080", nil) //如果Handle为nil则默认DefaultServeMux
	// http.ListenAndServe("localhost:8080", http.FileServer(http.Dir("assets")))
	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Header)
		fmt.Fprintln(w, r.Header["Accept-Encoding"])
		fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))
	})
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	})
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		id := params.Get("id")
		name := params.Get("name")
		log.Println(id)
		log.Println(name)
		log.Println(params)
	})
	http.ListenAndServe("localhost:8080", nil)
}
