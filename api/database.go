package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://Vipul1025:Vipul%401025@orderlocation.1gzh8eu.mongodb.net/?retryWrites=true&w=majority&appName=orderLocation")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	return client
}
