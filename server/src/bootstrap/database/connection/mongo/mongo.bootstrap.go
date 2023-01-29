package mongo

import (
	"context"
	"log"
	"os"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Conn *mongodriver.Database
}

var mongoSingleton *MongoConnection

func GetMongoConnection() *MongoConnection {
	if mongoSingleton == nil {
		uri := os.Getenv("MONGODB_URI")
		client, err := mongodriver.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		conn := client.Database(os.Getenv("MONGODB_DATABASE"))
		mongoSingleton = &MongoConnection{
			Conn: conn,
		}
		log.Println("Mongo connection created")
	}
	return mongoSingleton
}
