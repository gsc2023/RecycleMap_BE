package repository

import (
	"domain"
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
)

func CreateUser(user domain.User) (*firestore.DocumentRef, *firestore.WriteResult) {
	client := config.GetFirestore()

	ref, wr, err := client.Collection("users").Add(config.Ctx, user)
	if err != nil {
		log.Fatalf("error save user: %v\n", err)
	}

	defer client.Close()
	return ref, wr
}
