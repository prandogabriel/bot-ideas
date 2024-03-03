package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var cancel context.CancelFunc
var client *mongo.Client

func Connect() {
	ctx, c := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	cancel = c

	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	client = cl

	if err != nil {
		log.Fatal("error on connect database", err)
	}
}

func Cancel() {
	defer cancel()
}

func GetIdeas() {
	client.Ping(context.Background(), &readpref.ReadPref{})
}

func ADDIdea() {
	collection := client.Database("testing").Collection("ideas")
	res, _ := collection.InsertOne(context.Background(), bson.D{{"name", "pi"}, {"value", 3.14159}})
	id := res.InsertedID

	println(id)
}
