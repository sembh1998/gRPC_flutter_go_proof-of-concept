package mongo

import (
	"context"
	"log"
	"os"

	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/bootstrap"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/util"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Conn *mongodriver.Database
}

var mongoSingleton *MongoConnection

func NewMongoConnection() (bootstrap.Bootstrap, error) {
	return &MongoConnection{}, nil
}

func (m *MongoConnection) Initialice(config util.Config) error {
	uri := config.MongoDBURI
	client, err := mongodriver.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	conn := client.Database(config.MongoDatabase)
	mongoSingleton = &MongoConnection{
		Conn: conn,
	}
	log.Println("Mongo connection created")
	return nil
}

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
