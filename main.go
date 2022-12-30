package main

import (
	"github.com/lemoba/go-admin/controller"
	"github.com/lemoba/go-admin/middleware"
	"net/http"
)

func main() {
	controller.RegisterRoutes()
	http.ListenAndServe("localhost:8080", new(middleware.AuthMiddleware))
}
