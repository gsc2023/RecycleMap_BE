package service

import (
	"domain"
	"repository"

	"cloud.google.com/go/firestore"
)

func FindReports() []domain.ReportDto {
	return repository.FindAllReports()
}

func FindReport() {

}

func Join(report domain.Report) (*firestore.DocumentRef, *firestore.WriteResult) {
	return repository.SaveReport(report)
}

func DelReport() {

}

func ModifyReport() {

}

func ToggleLikeOfReport() {

}
