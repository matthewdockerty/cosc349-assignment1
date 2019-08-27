package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// const RESOURCE_FOLDER = "/vagrant/webapp/static"
const resourceFolder = "static"

type Recipe struct {
	Name        string
	Method      string
	Ingredients []string
	Image       []byte
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static/add.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form", 400)
			return
		}

		file, info, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Image must be provided", 500)
			return
		}
		contentType := info.Header.Get("Content-Type")
		if contentType != "image/jpeg" && contentType != "image/png" {
			http.Error(w, "Only JPG and PNG images are supported", 400)
			return
		}

		bs, err := ioutil.ReadAll(file)

		if err != nil {
			http.Error(w, "Unable to process image", 500)
			return
		}

		if len(bs) > 5e6 {
			http.Error(w, "Image too large (5MB limit)", 400)
			return
		}

		_, _, err = image.Decode(bytes.NewReader(bs))
		if err != nil {
			http.Error(w, "The provided image is not valid", 400)
			return
		}

		name := r.FormValue("name")
		method := r.FormValue("method")

		if len(name) == 0 || len(method) == 0 {
			http.Error(w, "Missing required fields", 400)
			return
		}

		i := 1
		var ingredients []string
		ingredient := r.FormValue(strings.Join([]string{"ingredient-", strconv.Itoa(i)}, ""))
		for len(ingredient) != 0 {
			ingredients = append(ingredients, ingredient)
			i++
			ingredient = r.FormValue(strings.Join([]string{"ingredient-", strconv.Itoa(i)}, ""))
		}

		if len(ingredients) < 1 {
			http.Error(w, "A recipe must contain at least one ingredient", 400)
			return
		}

		recipe := &Recipe{
			Name:        name,
			Method:      method,
			Ingredients: ingredients,
			Image:       bs,
		}

		fmt.Printf("%+v\n", recipe)
		// TODO: Store in db...

	default:
		http.Error(w, "Method Not Supported", 405)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		http.ServeFile(w, r, "static/index.html")
	case "/recipes", "/recipes/":
		http.ServeFile(w, r, "static/recipes.html")
	case "/add", "/add/":
		handleAdd(w, r)
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	http.HandleFunc("/", requestHandler)

	log.Println("Starting recipe webapp server " + os.Getenv("SERVER_NAME"))
	http.ListenAndServe(":3000", nil)
}
