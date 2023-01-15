package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Funca() {
	log.Printf("Funca server bootstrap")
}

type MongoConnection struct {
	Conn *mongo.Database
}

var mongoSingleton *MongoConnection

func GetMongoConnection() *MongoConnection {
	if mongoSingleton == nil {
		uri := os.Getenv("MONGODB_URI")
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		conn := client.Database("database01")
		mongoSingleton = &MongoConnection{
			Conn: conn,
		}
		log.Println("Mongo connection created")
	}
	return mongoSingleton
}
