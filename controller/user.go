package controller

import "net/http"

func registerUserRoutes() {
	http.HandleFunc("/login", login)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}
