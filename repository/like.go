package repository

import (
	"domain"
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

func SaveLike(like domain.LikeDao) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	ref, wr, err := config.GetFirestore().Collection("likes").Add(config.Ctx, like)
	if err != nil {
		log.Printf("error save like: %v\n", err)
	}

	return ref, wr, err
}

func FindLikeByID(ID string) (domain.LikeDto, error) {
	like := domain.LikeDao{}

	dsnap, err := config.GetFirestore().Collection("likes").Doc(ID).Get(config.Ctx)
	if err != nil {
		log.Printf("error find like by id: %v\n", err)
		return domain.LikeDto{ID: ID, Like: like}, err
	}

	err = mapstructure.Decode(dsnap.Data(), &like)
	if err != nil {
		log.Printf("error find like by id: %v\n", err)
		return domain.LikeDto{ID: ID, Like: like}, err
	}

	return domain.LikeDto{ID: ID, Like: like}, err
}

func FindLikeByUID(UID string) ([]domain.LikeDto, error) {
	likeDtos := []domain.LikeDto{}
	iter := config.GetFirestore().Collection("likes").Where("UID", "==", UID).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find likes by UID: %v\n", err)
			return likeDtos, err
		}

		like := domain.LikeDao{}
		err = mapstructure.Decode(doc.Data(), &like)
		if err != nil {
			log.Printf("error find likes by UID: %v\n", err)
			return likeDtos, err
		}

		likeDtos = append(likeDtos, domain.LikeDto{ID: doc.Ref.ID, Like: like})
	}

	return likeDtos, nil
}

func FindLikeByReportID(ReportID string) ([]domain.LikeDto, error) {
	likeDtos := []domain.LikeDto{}
	iter := config.GetFirestore().Collection("likes").Where("ReportID", "==", ReportID).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find likes by ReportID: %v\n", err)
			return likeDtos, err
		}

		like := domain.LikeDao{}
		err = mapstructure.Decode(doc.Data(), &like)
		if err != nil {
			log.Printf("error find likes by ReportID: %v\n", err)
			return likeDtos, err
		}

		likeDtos = append(likeDtos, domain.LikeDto{ID: doc.Ref.ID, Like: like})
	}

	return likeDtos, nil
}

func FindLikeByUIDAndReportID(UID string, ReportID string) ([]domain.LikeDto, error) {
	likeDtos := []domain.LikeDto{}
	iter := config.GetFirestore().Collection("likes").Where("UID", "==", UID).Where("ReportID", "==", ReportID).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find likes by UID, ReportID: %v\n", err)
			return likeDtos, err
		}

		like := domain.LikeDao{}
		err = mapstructure.Decode(doc.Data(), &like)
		if err != nil {
			log.Printf("error find likes by UID, ReportID: %v\n", err)
			return likeDtos, err
		}

		likeDtos = append(likeDtos, domain.LikeDto{ID: doc.Ref.ID, Like: like})
	}

	return likeDtos, nil
}

func FindAllLike() ([]domain.LikeDto, error) {
	likeDtos := []domain.LikeDto{}
	iter := config.GetFirestore().Collection("likes").Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find all likes: %v\n", err)
			return likeDtos, err
		}

		like := domain.LikeDao{}
		err = mapstructure.Decode(doc.Data(), &like)
		if err != nil {
			log.Printf("error find all likes: %v\n", err)
			return likeDtos, err
		}

		likeDtos = append(likeDtos, domain.LikeDto{ID: doc.Ref.ID, Like: like})
	}

	return likeDtos, nil
}

func SetLike(ID string, like domain.LikeDao) (*firestore.WriteResult, error) {
	wr, err := config.GetFirestore().Collection("likes").Doc(ID).Set(config.Ctx, like)
	if err != nil {
		log.Printf("error set like: %v\n", err)
	}

	return wr, err
}

func DelLike(ID string) (*firestore.WriteResult, error) {
	wr, err := config.GetFirestore().Collection("likes").Doc(ID).Delete(config.Ctx)
	if err != nil {
		log.Printf("error delete like: %v\n", err)
	}

	return wr, err
}
