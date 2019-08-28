package main

import (
	"bytes"
	"encoding/base64"
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
	"text/template"
)

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

		image := base64.StdEncoding.EncodeToString(bs)

		recipe := &Recipe{
			Name:        name,
			Method:      method,
			Ingredients: ingredients,
			Image:       image,
		}

		id, err := StoreRecipe(recipe)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/recipe?%s", id), http.StatusFound)

	default:
		http.Error(w, "Method Not Supported", 405)
	}
}

func handleRecipe(w http.ResponseWriter, r *http.Request) {
	recipeID := r.URL.RawQuery
	recipe, err := GetRecipeById(recipeID)

	if err != nil {
		http.Error(w, "Recipe not found", 404)
		return
	}

	t, err := template.ParseFiles("static/recipe.html")
	if err != nil {
		log.Print(err)
		return
	}
	t.Execute(w, recipe)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		http.ServeFile(w, r, "static/index.html")
	case "/recipes", "/recipes/":
		http.ServeFile(w, r, "static/recipes.html")
	case "/add", "/add/":
		handleAdd(w, r)
	case "/background.jpg":
		http.ServeFile(w, r, "static/background.jpg")
	case "/recipe":
		handleRecipe(w, r)
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	log.Println("Connecting to MongoDB...")
	if err := InitDB(); err != nil {
		log.Panic("Unable to connect to database")
	}

	log.Println("Starting recipe webapp server " + os.Getenv("SERVER_NAME"))
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":3000", nil)
}
