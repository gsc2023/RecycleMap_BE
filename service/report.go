package service

import (
	"domain"
	"repository"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
)

func FindReports() ([]domain.ReportDto, error) {
	return repository.FindAllReports()
}

func FindReport(ID string) (domain.ReportDto, error) {
	return repository.FindReportByID(ID)
}

func JoinReport(token *auth.Token, report domain.ReportDao) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return repository.SaveReport(token.UID, report)
}

func DelReport(ID string) error {
	return repository.DelReport(ID)
}

func ModifyReport(ID string, report domain.ReportDao) (*firestore.WriteResult, error) {
	return repository.SetReport(ID, report)
}

func ToggleLikeOfReport(token *auth.Token, ID string) (status bool, err error) {
	return
}
