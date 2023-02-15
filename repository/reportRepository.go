package repository

import (
	"domain"
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

func SaveReport(report domain.Report) (*firestore.DocumentRef, *firestore.WriteResult) {
	ref, wr, err := config.GetFirestore().Collection("reports").Add(config.Ctx, report)
	if err != nil {
		log.Fatalf("error save report: %v\n", err)
	}

	return ref, wr
}

func FindReportByID(ID string) domain.ReportDto {
	dsnap, err := config.GetFirestore().Collection("reports").Doc(ID).Get(config.Ctx)
	if err != nil {
		log.Fatalf("error find report by id: %v\n", err)
	}

	report := domain.Report{}
	err = mapstructure.Decode(dsnap.Data(), &report)
	if err != nil {
		log.Fatalf("error find report by id: %v\n", err)
	}

	return domain.ReportDto{ID: ID, Report: report}
}

func FindAllReports() (reportDtos []domain.ReportDto) {
	iter := config.GetFirestore().Collection("reports").Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error find all reports: %v\n", err)
		}

		report := domain.Report{}
		err = mapstructure.Decode(doc.Data(), &report)
		if err != nil {
			log.Fatalf("error find all reports: %v\n", err)
		}

		reportDtos = append(reportDtos, domain.ReportDto{ID: doc.Ref.ID, Report: report})
	}

	return
}

func SetReport(ID string, report *domain.Report) *firestore.WriteResult {
	wr, err := config.GetFirestore().Collection("reports").Doc(ID).Set(config.Ctx, report)
	if err != nil {
		log.Fatalf("error set report: %v\n", err)
	}

	return wr
}

func FindReportsByUserId(userID string) (reportDtos []domain.ReportDto) {
	iter := config.GetFirestore().Collection("reports").Where("userID", "==", userID).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error find reports by userID: %v\n", err)
		}

		report := domain.Report{}
		err = mapstructure.Decode(doc.Data(), &report)
		if err != nil {
			log.Fatalf("error find all reports: %v\n", err)
		}

		reportDtos = append(reportDtos, domain.ReportDto{ID: doc.Ref.ID, Report: report})
	}

	return
}
