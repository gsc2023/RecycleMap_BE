package service

import (
	"domain"
	"repository"

	"cloud.google.com/go/firestore"
)

func FindReports() []domain.ReportDto {
	return repository.FindAllReports()
}

func FindReport(ID string) domain.ReportDto {
	return repository.FindReportByID(ID)
}

func JoinReport(report domain.Report) (*firestore.DocumentRef, *firestore.WriteResult) {
	return repository.SaveReport(report)
}

func DelReport(ID string) {
	repository.DelReport(ID)
}

func ModifyReport(ID string, report domain.Report) {
	repository.SetReport(ID, report)
}

func ToggleLikeOfReport() {

}
