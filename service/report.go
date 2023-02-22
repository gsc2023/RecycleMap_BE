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
	likes, err := repository.FindLikeByUIDAndReportID(token.UID, ID)

	if err != nil {
		log.Printf("error toggle like: %v\n", err)
		return false, err
	}

	_, err = repository.FindReportByID(ID)

	if err != nil {
		log.Printf("error toggle like: %v\n", err)
		return false, err
	}

	if len(likes) == 0 { // 저장 안되어있음
		like := domain.LikeDao{UID: token.UID, ReportID: ID, Status: true}
		_, _, err := repository.SaveLike(like)

		if err != nil {
			log.Printf("error toggle like: %v\n", err)
			return false, err
		}

		err = HandleLikeOfReport(ID, true)

		if err != nil {
			log.Printf("error toggle like: %v\n", err)
			return false, err
		}

		return true, err
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

	err = HandleLikeOfReport(ID, toggleLike) // report 좋아요 조절하기

	if err != nil {
		log.Printf("error toggle like: %v\n", err)
		return false, err
	}

	return toggleLike, err
}

func HandleLikeOfReport(ID string, status bool) error {
	reportDto, err := repository.FindReportByID(ID)

	if err != nil {
		log.Printf("error handle like: %v\n", err)
		return err
	}

	if status {
		reportDto.Report.Like++
	} else {
		reportDto.Report.Like--
	}

	_, err = repository.SetReport(ID, reportDto.Report)

	if err != nil {
		log.Printf("error handle like: %v\n", err)
	}

	return err
}
