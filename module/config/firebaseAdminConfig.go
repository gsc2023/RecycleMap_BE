package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var Ctx = context.Background()

func InitializeAppWithServiceAccount() *firebase.App {
	// [START initialize_app_service_account_golang]

	sa := option.WithCredentialsFile("module/config/serviceAccountKey.json")

	app, err := firebase.NewApp(Ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_app_service_account_golang]

	return app
}
