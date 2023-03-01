package service

import (
	"domain"
	"repository"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
)

func JoinComment(token *auth.Token, ID string, comment domain.Comment) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return repository.SaveComment(token.UID, ID, comment)
}

func FindCommentsById(ID string) ([]domain.CommentDto, error) {
	return repository.FindAllCommentsById(ID)
}

func FindCommentsByUID(token *auth.Token) ([]domain.CommentDto, error) {
	return repository.FindAllcommentsByUID(token.UID)
}
