package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets"+r.URL.Path)
	})

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

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// r.ParseForm() // 如果表单和URL里有同样的key， 那么它们都会放在一个slice里：表单里的值靠前，URL的值靠后
		//r.ParseMultipartForm(1024) // length为字节数
		//fmt.Fprintln(w, r.Form)
		//fmt.Fprintln(w, r.PostForm) // 只获取表单的key-value对
		//fmt.Fprintln(w, r.Form["first_name"])
		//fmt.Fprintln(w, r.FormValue("first_name")) // 只取表单的值 FormValue和PostFormValue都会调用ParseMultipartForm
		//fmt.Fprintln(w, r.PostFormValue("first_name"))
		//fileHeader := r.MultipartForm.File["uploaded"][0]
		//file, err := fileHeader.Open()

		file, _, err := r.FormFile("uploaded")
		if err == nil {
			data, err := io.ReadAll(file)
			if err == nil {
				fmt.Fprintln(w, string(data))
			}
		}
	})

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	http.ListenAndServe("localhost:8080", nil)
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html><head>Go Web Programming!</head></html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No Such Service")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://baidu.com")
	w.WriteHeader(302)
}

type Post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "ranen",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}
