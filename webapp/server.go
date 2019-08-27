package main

import (
	"log"
	"net/http"
	"os"
)

// const RESOURCE_FOLDER = "/vagrant/webapp/static"
const resourceFolder = "static"

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		http.ServeFile(w, r, "static/index.html")
	case "/recipes", "/recipes/":
		http.ServeFile(w, r, "static/recipes.html")
	case "/add", "/add/":
		http.ServeFile(w, r, "static/add.html")
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	http.HandleFunc("/", requestHandler)

	log.Println("Starting recipe webapp server " + os.Getenv("SERVER_NAME"))
	http.ListenAndServe(":3000", nil)
}
