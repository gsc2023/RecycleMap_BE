package service

import (
	"domain"
	"log"
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

func ModifyComment(token *auth.Token, ID string, newComment domain.Comment) (*firestore.WriteResult, error) {
	err := IsOwner(repository.IsCommentOwner(token.UID, ID))
	if err != nil {
		log.Printf("error modify comment: %v\n", err)
		return nil, err
	}

	commentDto, err := repository.FindCommentById(ID)
	if err != nil {
		log.Printf("error modify comment: %v\n", err)
		return nil, err
	}

	newComment.UID = token.UID
	newComment.LocationID = commentDto.Comment.LocationID

	return repository.SetComment(ID, newComment)
}

func DeleteComment(token *auth.Token, ID string) (*firestore.WriteResult, error) {
	err := IsOwner(repository.IsCommentOwner(token.UID, ID))

	if err != nil {
		log.Printf("error delete comment: %v\n", err)
		return nil, err
	}

	return repository.DeleteComment(ID)
}
