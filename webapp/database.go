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

type Recipe struct {
	Name        string
	Method      string
	Ingredients []string
	Image       string // base64 encoded
}

var db *mongo.Database

var dbHost = os.Getenv("DB_HOST")

const dbName = "recipesdb"
const recipesCollectionName = "recipes"

var recipesCollection *mongo.Collection

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

func StoreRecipe(r *Recipe) (string, error) {
	res, err := recipesCollection.InsertOne(context.TODO(), r)

	if err != nil {
		return "", errors.New("Unable to insert recipe into collection")
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

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

func DeleteRecipeByID(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := recipesCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})

	if err != nil {
		return errors.New("Unable to delete recipe")
	}

	return nil
}
