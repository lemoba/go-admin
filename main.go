package main

import (
	"github.com/lemoba/go-admin/controller"
	"net/http"
)

func main() {
	controller.RegisterRoutes()
	http.ListenAndServe("localhost:8080", nil)
}
