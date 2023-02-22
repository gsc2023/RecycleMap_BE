package repository

import (
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
)

func SaveComment(ID string, content string, uID string) (*firestore.DocumentRef, *firestore.WriteResult) {
	ref, wr, err := config.GetFirestore().Collection("comments").Add(config.Ctx, comment)
	if err != nil {
		log.Printf("error save comment: %v\n", err)
	}

	return ref, wr
}
