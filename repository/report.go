package repository

import (
	"domain"
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

func SaveReport(UID string, report domain.ReportDao) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	report.UID = UID
	ref, wr, err := config.GetFirestore().Collection("reports").Add(config.Ctx, report)
	if err != nil {
		log.Printf("error save report: %v\n", err)
	}

	return ref, wr, err
}

func FindReportByID(ID string) (domain.ReportDto, error) {
	report := domain.ReportDao{}

	dsnap, err := config.GetFirestore().Collection("reports").Doc(ID).Get(config.Ctx)
	if err != nil {
		log.Printf("error find report by id: %v\n", err)
		return domain.ReportDto{ID: ID, Report: report}, err
	}

	err = mapstructure.Decode(dsnap.Data(), &report)
	if err != nil {
		log.Printf("error find report by id: %v\n", err)
		return domain.ReportDto{ID: ID, Report: report}, err
	}

	return domain.ReportDto{ID: ID, Report: report}, err
}

func FindAllReports() ([]domain.ReportDto, error) {
	reportDtos := []domain.ReportDto{}
	iter := config.GetFirestore().Collection("reports").Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find all reports: %v\n", err)
			return reportDtos, err
		}

		report := domain.ReportDao{}
		err = mapstructure.Decode(doc.Data(), &report)
		if err != nil {
			log.Printf("error find all reports: %v\n", err)
			return reportDtos, err
		}

		reportDtos = append(reportDtos, domain.ReportDto{ID: doc.Ref.ID, Report: report})
	}

	return reportDtos, nil
}

func SetReport(ID string, report domain.ReportDao) (*firestore.WriteResult, error) {
	wr, err := config.GetFirestore().Collection("reports").Doc(ID).Set(config.Ctx, report)
	if err != nil {
		log.Printf("error set report: %v\n", err)
	}

	return wr, err
}

func FindReportsByUId(UID string) ([]domain.ReportDto, error) {
	reportDtos := []domain.ReportDto{}
	iter := config.GetFirestore().Collection("reports").Where("UID", "==", UID).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find reports by UID: %v\n", err)
			return reportDtos, err
		}

		report := domain.ReportDao{}
		err = mapstructure.Decode(doc.Data(), &report)
		if err != nil {
			log.Printf("error find reports by UID: %v\n", err)
			return reportDtos, err
		}

		reportDtos = append(reportDtos, domain.ReportDto{ID: doc.Ref.ID, Report: report})
	}

	return reportDtos, nil
}

func DelReport(ID string) (*firestore.WriteResult, error) {
	wr, err := config.GetFirestore().Collection("reports").Doc(ID).Delete(config.Ctx)
	if err != nil {
		log.Printf("error delete report: %v\n", err)
	}
	return wr, err
}
