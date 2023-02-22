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
	return repository.FindReportsExecptDisabled()
}

func FindReport(ID string) (domain.ReportDto, error) {
	return repository.FindReportByID(ID)
}

func JoinReport(token *auth.Token, report domain.ReportDao) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return repository.SaveReport(token.UID, report)
}

func DelReport(token *auth.Token, ID string) (*firestore.WriteResult, error) {
	err := IsOwner(repository.IsReportOwner(token.UID, ID))

	if err != nil {
		log.Printf("error delete report: %v\n", err)
		return nil, err
	}

	err = delLike(ID)

	if err != nil {
		log.Printf("error delete report: %v\n", err)
		return nil, err
	}

	return repository.DelReport(ID)
}

func ModifyReport(token *auth.Token, ID string, newReport domain.ReportDao) (*firestore.WriteResult, error) {
	err := IsOwner(repository.IsReportOwner(token.UID, ID))

	if err != nil {
		log.Printf("error modify report: %v\n", err)
		return nil, err
	}

	reportDto, err := repository.FindReportByID(ID)

	if err != nil {
		log.Printf("error modify report: %v\n", err)
		return nil, err
	}

	newReport.UID = token.UID
	newReport.Like = reportDto.Report.Like
	newReport.Disabled = reportDto.Report.Disabled

	return repository.SetReport(ID, newReport)
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

	cnt, err := handleLikeOfReport(ID, toggleLike) // report 좋아요 조절하기

	if err != nil {
		log.Printf("error toggle like: %v\n", err)
		return false, err
	}

	if cnt >= LIKE_WHEN_BE_LOCATION {
		location := reportDtoToLocation(reportDto)
		makeReportDisable(ID)
		SaveLocation(location) // todo: error 처리
	}

	return toggleLike, err
}

func handleLikeOfReport(ID string, status bool) (int, error) {
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

func reportDtoToLocation(report domain.ReportDto) (location domain.Location) {
	location.Name = report.Report.Name
	location.LocationType = report.Report.LocationType
	location.Longitude = report.Report.Longitude
	location.Latitude = report.Report.Latitude
	location.Content = report.Report.Content

	return
}

func makeReportDisable(ID string) error {
	reportDto, err := repository.FindReportByID(ID)

	if err != nil {
		log.Printf("error handle like: %v\n", err)
		return err
	}

	reportDto.Report.Disabled = true

	_, err = repository.SetReport(ID, reportDto.Report)

	if err != nil {
		log.Printf("error handle like: %v\n", err)
	}

	return err
}

func delLike(ID string) error {
	likeDtos, err := repository.FindLikeByReportID(ID)

	if err != nil {
		log.Printf("error delete like: %v\n", err)
		return err
	}

	for _, val := range likeDtos {
		_, err := repository.DelLike(val.ID)

		if err != nil {
			log.Printf("error delete like: %v\n", err)
		}
	}

	return nil
}
