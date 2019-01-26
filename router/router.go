package main

import (
	"log"
	"net/http"
	"time"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/", requestHandler)

	server := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
