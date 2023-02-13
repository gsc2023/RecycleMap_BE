package repository

import (
	"domain"
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
)

func SaveReport(report domain.Report) (*firestore.DocumentRef, *firestore.WriteResult) {
	client := config.GetFirestore()

	ref, wr, err := client.Collection("reports").Add(config.Ctx, report)
	if err != nil {
		log.Fatalf("error save report: %v\n", err)
	}

	defer client.Close()
	return ref, wr
}

func FindReportByID() (ID string, report *domain.Report) {
	client := config.GetFirestore()

	dsnap, err := client.Collection("reports").Doc(ID).Get(config.Ctx)
	if err != nil {
		log.Fatalf("error find report by id: %v\n", err)
	}

	err = mapstructure.Decode(dsnap.Data(), &report)
	if err != nil {
		log.Fatalf("error find report by id: %v\n", err)
	}

	defer client.Close()
	return
}

func FindReportsByUserId() {

}

func FindReportByName() {

}

func FindAllReports() (reports []domain.Report) {
	client := config.GetFirestore()

	dsnap, err := client.Collection("reports").Documents(config.Ctx).GetAll()
	if err != nil {
		log.Fatalf("error find all reports: %v\n", err)
	}

	for _, val := range dsnap {
		report := domain.Report{}
		err = mapstructure.Decode(val.Data(), &report)
		if err != nil {
			log.Fatalf("error find all reports: %v\n", err)
		}

		log.Println(val.Data(), report)

		reports = append(reports, report)
	}

	defer client.Close()
	return
}

func ModifyReport(ID string, report *domain.Report) *firestore.WriteResult {
	client := config.GetFirestore()

	wr, err := client.Collection("reports").Doc(ID).Set(config.Ctx, report)
	if err != nil {
		log.Fatalf("error modify report: %v\n", err)
	}

	defer client.Close()
	return wr
}
