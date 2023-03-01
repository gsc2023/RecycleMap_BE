package repository

import (
	"domain"
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
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

func FindAllCommentsById(ID string) ([]domain.CommentDto, error) {
	commentDtos := []domain.CommentDto{}
	iter := config.GetFirestore().Collection("comments").Where("LocationID", "==", ID).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find all comments by locationId: %v\n", err)
			return commentDtos, err
		}

		comment := domain.Comment{}
		err = mapstructure.Decode(doc.Data(), &comment)
		if err != nil {
			log.Printf("error find all comments by locationId: %v\n", err)
			return commentDtos, err
		}

		commentDtos = append(commentDtos, domain.CommentDto{ID: doc.Ref.ID, Comment: comment})
	}

	return commentDtos, nil
}

func FindAllcommentsByUID(UID string) ([]domain.CommentDto, error) {
	commentDtos := []domain.CommentDto{}
	iter := config.GetFirestore().Collection("comments").Where("UID", "==", UID).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find all comments by locationId: %v\n", err)
			return commentDtos, err
		}

		comment := domain.Comment{}
		err = mapstructure.Decode(doc.Data(), &comment)
		if err != nil {
			log.Printf("error find all comments by locationId: %v\n", err)
			return commentDtos, err
		}

		commentDtos = append(commentDtos, domain.CommentDto{ID: doc.Ref.ID, Comment: comment})
	}

	return commentDtos, nil
}
