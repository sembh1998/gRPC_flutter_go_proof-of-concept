package firebase

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

type FirebaseConnection struct {
	Conn *db.Client
}

var firebaseSingleton *FirebaseConnection

func GetMongoConnection() *FirebaseConnection {
	if firebaseSingleton == nil {
		ctx := context.Background()

		// configure database URL
		conf := &firebase.Config{
			DatabaseURL: os.Getenv("FIREBASE_DATABASE_URL"), //FIREBASE_DATABASE_URL
		}

		// fetch service account key
		opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICE_ACCOUNT_KEY_JSON_PATH")) //FIREBASE_SERVICE_ACCOUNT_KEY_JSON_PATH

		app, err := firebase.NewApp(ctx, conf, opt)
		if err != nil {
			panic(err)
		}

		client, err := app.Database(ctx)
		if err != nil {
			panic(err)
		}
		firebaseSingleton = &FirebaseConnection{
			Conn: client,
		}

	}
	return firebaseSingleton
}
