package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("/vagrant/webapp/static")).ServeHTTP(w, r)
	})

	log.Println("Starting recipe webapp server " + os.Getenv("SERVER_NAME"))
	http.ListenAndServe(":3000", nil)
}
