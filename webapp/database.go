package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type key string

type Recipe struct {
	Name        string
	Method      string
	Ingredients []string
	Image       []byte
}

var db *mongo.Database
var ctx context.Context

func InitDB() error {
	ctx = context.Background()
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
	res, err := db.Collection("recipes").InsertOne(ctx, r)

	if err != nil {
		return "", errors.New("Unable to insert recipe into collection")
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
