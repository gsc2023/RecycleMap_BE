package repository

import (
	"domain"
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
)

func SaveComment(UID string, ID string, comment domain.Comment) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	comment.UID = UID
	comment.LocationID = ID

	ref, wr, err := config.GetFirestore().Collection("comments").Add(config.Ctx, comment)
	if err != nil {
		log.Printf("error SaveComment: %v\n", err)
	}

	return ref, wr, err
}
