package main

import (
	"fmt"
	"net/http"
	"work-manager/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 3001),
		Handler: router,
		// ReadTimeout:    setting.ReadTimeout,
		// WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
