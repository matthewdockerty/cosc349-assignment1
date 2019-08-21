package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world from: "+os.Getenv("SERVER_NAME"))
	})

	log.Println("Starting recipe webapp server " + os.Getenv("SERVER_NAME"))
	http.ListenAndServe(":3000", nil)
}
