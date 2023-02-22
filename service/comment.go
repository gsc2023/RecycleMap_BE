package service

import (
	"repository"

	"cloud.google.com/go/firestore"
)

func SaveComment(ID string, content string, uID string) (*firestore.DocumentRef, *firestore.WriteResult) {
	return repository.SaveComment(ID, content, uID)
}
