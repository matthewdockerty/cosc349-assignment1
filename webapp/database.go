package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Recipe struct stores data for a recipe.
type Recipe struct {
	Name        string
	Method      string
	Ingredients []string
	Image       string // base64 encoded
}

// Database host address & port, database name, & collection name constants
var dbHost = os.Getenv("DB_HOST") // Environment variable is set by Vagrant during the provisioning process
const dbName = "recipesdb"
const recipesCollectionName = "recipes"

// Pointers to database & collection structs.
var db *mongo.Database
var recipesCollection *mongo.Collection

/*
InitDB connects to the MongoDB database and initializes the values of
the database & collection pointers defined above.

This function returns an error if the connection was unsuccessful.
*/
func InitDB() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	uri := fmt.Sprintf(`mongodb://%s/%s`, dbHost, dbName)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return errors.New("Unable to connect to MongoDB")
	}
	err = client.Connect(ctx)
	if err != nil {
		return errors.New("Mongo client couldn't connect with background context")
	}

	db = client.Database(dbName)
	recipesCollection = db.Collection(recipesCollectionName)

	return nil
}

/*
StoreRecipe stores the data for the given recipe in the database.
No input checking or serialization is performed here.

It returns the object ID of the recipe (or an empty string if storing
was unsuccessful) and an error if one was encountered while attempting
storage.
*/
func StoreRecipe(r *Recipe) (string, error) {
	res, err := recipesCollection.InsertOne(context.TODO(), r)

	if err != nil {
		return "", errors.New("Unable to insert recipe into collection")
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

/*
GetRecipeByID retrieves the recipe with the given object ID from the
database and an error if any are encountered (e.g. if no recipe is
associated with the ID).
*/
func GetRecipeByID(id string) (Recipe, error) {
	var result Recipe
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	err := recipesCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, errors.New("Recipe not found")
	}

	return result, nil
}

/*
GetAllRecipes retrieves all recipes from the database and returns them in a
map with object IDs as keys and recipe structs as values. It also returns
any errors encountered.
*/
func GetAllRecipes() (map[string]Recipe, error) {

	cursor, err := recipesCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, errors.New("Unable to get all recipes")
	}
	defer cursor.Close(context.TODO())

	results := make(map[string]Recipe)

	for cursor.Next(context.TODO()) {
		var r Recipe
		err := cursor.Decode(&r)
		if err != nil {
			return nil, errors.New("Unable to decode recipe")
		}

		// Get object ID from cursor
		vals, _ := cursor.Current.Values()
		var id primitive.ObjectID
		vals[0].Unmarshal(&id)

		results[id.Hex()] = r
	}

	return results, nil
}

/*
DeleteRecipeByID deletes the recipe with the given object ID from the
database. It returns an error if one was encountered.
*/
func DeleteRecipeByID(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := recipesCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})

	if err != nil {
		return errors.New("Unable to delete recipe")
	}

	return nil
}
