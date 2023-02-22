package service

import (
	"domain"
	"log"
	"repository"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
)

const LikeWhenReportToLocation = 50

func FindReports() ([]domain.ReportDto, error) {
	return repository.FindAllReports()
}

func FindReport(ID string) (domain.ReportDto, error) {
	return repository.FindReportByID(ID)
}

func JoinReport(token *auth.Token, report domain.ReportDao) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return repository.SaveReport(token.UID, report)
}

func DelReport(ID string) (*firestore.WriteResult, error) {
	return repository.DelReport(ID)
}

func ModifyReport(ID string, report domain.ReportDao) (*firestore.WriteResult, error) {
	return repository.SetReport(ID, report)
}

func ToggleLikeOfReport(token *auth.Token, ID string) (status bool, err error) {
	likes, err := repository.FindLikeByUIDAndLocationID(token.UID, ID)

	if len(likes) == 0 { // 저장 안되어있음
		like := domain.LikeDao{UID: token.UID, ReportID: ID, Status: true}
		_, _, err := repository.SaveLike(like)

		if err != nil {
			log.Printf("error toggle like: %v\n", err)
			return false, err
		}

		return true, err
	}

	if err != nil {
		log.Printf("error toggle like: %v\n", err)
		return false, err
	}

	toggleLike := true

	if likes[0].Like.Status {
		toggleLike = false
	}

	like := domain.LikeDao{UID: token.UID, ReportID: ID, Status: toggleLike}

	_, err = repository.SetLike(likes[0].ID, like)

	if err != nil {
		log.Printf("error toggle like: %v\n", err)
		return false, err
	}

	return toggleLike, err
}
