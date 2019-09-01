package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html"
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

// Static content directory path. Environment variable set by Vagrant during provisioning.
var resourceDir = os.Getenv("CONTENT_PATH")

/* Server instance name. Each running instance has a different environment
variable value set during provisioning. This name is shown on the main recipes
page to demonstrate load balancing behaviour and to show which instance served
the request.
*/
var serverName = os.Getenv("SERVER_NAME")

/*
Response struct is used to pass the server instance name and a response of any
type to the template engine whenever we want to display the server name to the
user on the webpage.
*/
type response struct {
	ServerName string
	Response   interface{}
}

/*
handleAdd deals with all requests related to adding recipes.
URI: /add

It serves the add recipes form page and handles post requests from the
submitted form to save new recipes.
*/
func handleAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, resourceDir+"add.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form", 400)
			return
		}

		// Handle image upload. Read data & convert to base64

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

		if len(bs) > 1e6 {
			http.Error(w, "Image too large (1MB limit)", 400)
			return
		}

		_, _, err = image.Decode(bytes.NewReader(bs))
		if err != nil {
			http.Error(w, "The provided image is not valid", 400)
			return
		}

		image := base64.StdEncoding.EncodeToString(bs)

		// Read & escape all other fields from form data.

		name := r.FormValue("name")
		method := r.FormValue("method")

		name = html.EscapeString(name)
		method = html.EscapeString(method)
		method = strings.ReplaceAll(method, "\n", "<br>")

		if len(name) == 0 || len(method) == 0 {
			http.Error(w, "Missing required fields", 400)
			return
		}

		/* Read ingredients from form data - needs to be handled seperately
		as there can be an arbitrary number of ingredients */
		i := 1
		var ingredients []string
		ingredient := r.FormValue(strings.Join([]string{"ingredient-", strconv.Itoa(i)}, ""))
		ingredient = html.EscapeString(ingredient)
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
			Image:       image,
		}

		id, err := StoreRecipe(recipe)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Redirect to main recipes list page.
		http.Redirect(w, r, fmt.Sprintf("/recipe?%s", id), http.StatusFound)

	default:
		http.Error(w, "Method Not Supported", 405)
	}
}

/*
handleRecipe deals with all requests to view an individual recipe.
URI: /recipe?<recipe_id>
*/
func handleRecipe(w http.ResponseWriter, r *http.Request) {
	recipeID := r.URL.RawQuery
	recipe, err := GetRecipeByID(recipeID)

	if err != nil {
		http.Error(w, "Recipe not found", 404)
		return
	}

	t, err := template.ParseFiles(resourceDir + "recipe.html")
	if err != nil {
		http.Error(w, "Unable to parse template file", 500)
		return
	}
	t.Execute(w, recipe)
}

/*
handleRecipes deals with requests to view all recipes.
URI: /recipes
*/
func handleRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := GetAllRecipes()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.ParseFiles(resourceDir + "recipes.html")
	if err != nil {
		http.Error(w, "Unable to parse template file", 500)
		return
	}

	t.Execute(w, response{serverName, recipes})
}

/*
handleDelete deals with requests to delete a given recipe.
URI: /delete?<recipe_id>

This function accepts any HTTP method because the frontend does
not use AJAX requests, so we can't easily perform a DELETE request.
*/
func handleDelete(w http.ResponseWriter, r *http.Request) {
	recipeID := r.URL.RawQuery
	err := DeleteRecipeByID(recipeID)

	if err != nil {
		http.Error(w, "Unable to delete recipe", 500)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/recipes"), http.StatusMovedPermanently)
}

/*
	requestHandler handles all incoming requests and routes them to the correct
	function accordingly, or returns a 404 if a non-existent resource is requested.
*/
func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		http.ServeFile(w, r, resourceDir+"index.html")
	case "/recipes", "/recipes/":
		handleRecipes(w, r)
	case "/add", "/add/":
		handleAdd(w, r)
	case "/delete":
		handleDelete(w, r)
	case "/background.jpg":
		http.ServeFile(w, r, resourceDir+"background.jpg")
	case "/recipe":
		handleRecipe(w, r)
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

// main function initializes the database connection and starts the http server.
func main() {
	log.Println("Connecting to MongoDB...")
	if err := InitDB(); err != nil {
		log.Panic("Unable to connect to database")
	}

	log.Println("Starting recipe webapp server " + serverName + "...")
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":3000", nil)
}
