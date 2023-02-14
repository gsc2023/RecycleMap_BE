package service

import (
	"domain"
	"repository"

	"cloud.google.com/go/firestore"
)

func Signup(user domain.User) (*firestore.DocumentRef, *firestore.WriteResult) {
	return repository.CreateUser(user)
}
