package service

import (
	"log"
	"module/config"
)

func Example() {
	// Use a service account
	app := config.InitializeAppWithServiceAccount()

	client, err := app.Firestore(config.Ctx)
	if err != nil {
		log.Println("firestore")
		log.Fatalln(err)
	}

	_, _, err = client.Collection("users").Add(config.Ctx, map[string]interface{}{
		"first": "aaaaaaa",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	defer client.Close()
}
