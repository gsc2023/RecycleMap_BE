package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var Ctx = context.Background()
var App *firebase.App = initializeAppWithServiceAccount()

func initializeAppWithServiceAccount() *firebase.App {
	// [START initialize_app_service_account_golang]

	sa := option.WithCredentialsFile("module/config/serviceAccountKey.json")

	app, err := firebase.NewApp(Ctx, nil, sa)
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
