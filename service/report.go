package service

import (
	"domain"
	"log"
	"repository"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
)

const LIKE_WHEN_BE_LOCATION = 50

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

	reportDto, err := repository.FindReportByID(ID)

	if err != nil {
		log.Printf("error toggle like: %v\n", err)
		return false, err
	}

	var likeID string
	toggleLike := true

	if len(likes) != 0 {
		toggleLike = !likes[0].Like.Status
		likeID = likes[0].ID
	} else { // 저장 안되어있음
		ref, _, err := repository.SaveLike(domain.LikeDao{})

		if err != nil {
			log.Printf("error toggle like: %v\n", err)
			return false, err
		}

		likeID = ref.ID
	}

	like := domain.LikeDao{UID: token.UID, ReportID: ID, Status: toggleLike}

	_, err = repository.SetLike(likeID, like)

	if err != nil {
		log.Printf("error toggle like: %v\n", err)
		return false, err
	}

	cnt, err := HandleLikeOfReport(ID, toggleLike) // report 좋아요 조절하기

	if err != nil {
		log.Printf("error toggle like: %v\n", err)
		return false, err
	}

	if cnt >= LIKE_WHEN_BE_LOCATION {
		location := ReportDtoToLocation(reportDto)
		SaveLocation(location) // todo: error 처리
	}

	return toggleLike, err
}

func HandleLikeOfReport(ID string, status bool) (int, error) {
	reportDto, err := repository.FindReportByID(ID)

	if err != nil {
		log.Printf("error handle like: %v\n", err)
		return reportDto.Report.Like, err
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

	return reportDto.Report.Like, err
}

func ReportDtoToLocation(report domain.ReportDto) (location domain.Location) {
	location.Name = report.Report.Name
	location.LocationType = report.Report.LocationType
	location.Longitude = report.Report.Longitude
	location.Latitude = report.Report.Latitude
	location.Content = report.Report.Content
}
