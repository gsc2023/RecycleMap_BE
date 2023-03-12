package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

var Ctx = context.Background()
var App *firebase.App = initializeAppWithServiceAccount()

func initializeAppWithServiceAccount() *firebase.App {
	// [START initialize_app_service_account_golang]

	sa := option.WithCredentialsFile("module/config/serviceAccountKey.json")

	config := &firebase.Config{
		StorageBucket: "solutionchallenge2023.appspot.com",
	}

	app, err := firebase.NewApp(Ctx, config, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_app_service_account_golang]

	return app
}

func GetFirestore() *firestore.Client {
	client, err := App.Firestore(Ctx)
	if err != nil {
		log.Println("firestore")
		log.Fatalln(err)
	}

	return client
}

func GetAuth() *auth.Client {
	client, err := App.Auth(Ctx)
	if err != nil {
		log.Println("firestore")
		log.Fatalln(err)
	}

	return client
}

func GetStorage() *storage.Client {
	client, err := App.Storage(Ctx)
	if err != nil {
		log.Println("firestore")
		log.Fatalln(err)
	}

	return client
}
