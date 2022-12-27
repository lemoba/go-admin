package main

import (
	"fmt"
	"github.com/lemoba/go-admin/pkg/setting"
	"github.com/lemoba/go-admin/router"
	"net/http"
)

func main() {
	router := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
