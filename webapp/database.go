package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type key string

type Recipe struct {
	Name        string
	Method      string
	Ingredients []string
	Image       string // base64 encoded
}

var db *mongo.Database

func InitDB() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx = context.WithValue(ctx, key("host"), "localhost:27017")
	ctx = context.WithValue(ctx, key("database"), "recipesdb")

	uri := fmt.Sprintf(`mongodb://%s/%s`, ctx.Value(key("host")), ctx.Value(key("database")))

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return errors.New("Unable to connect to MongoDB")
	}
	err = client.Connect(ctx)
	if err != nil {
		return errors.New("Mongo client couldn't connect with background context")
	}

	db = client.Database("recipesdb")

	return nil
}

func StoreRecipe(r *Recipe) (string, error) {
	res, err := db.Collection("recipes").InsertOne(context.TODO(), r)

	if err != nil {
		return "", errors.New("Unable to insert recipe into collection")
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func GetRecipeById(id string) (Recipe, error) {
	var result Recipe
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	err := db.Collection("recipes").FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, errors.New("Recipe not found")
	}

	return result, nil
}

func GetAllRecipes() (map[string]Recipe, error) {

	cursor, err := db.Collection("recipes").Find(context.TODO(), bson.D{})
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
