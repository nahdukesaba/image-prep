package main

import (
	"image-prep/api"
	"log"
	"net/http"
	"time"
)

func main() {
	router := api.Router()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
