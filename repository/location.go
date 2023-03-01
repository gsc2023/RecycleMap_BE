package repository

import (
	"domain"
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

func FindAllLocations() ([]domain.LocationDto, error) {
	locationDtos := []domain.LocationDto{}
	iter := config.GetFirestore().Collection("locations").Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find all locations: %v\n", err)
			return locationDtos, err
		}

		location := domain.Location{}
		err = mapstructure.Decode(doc.Data(), &location)
		if err != nil {
			log.Printf("error find all locations: %v\n", err)
			return locationDtos, err
		}

		locationDtos = append(locationDtos, domain.LocationDto{ID: doc.Ref.ID, Location: location})
	}

	return locationDtos, nil
}

func FindLocationById(ID string) (domain.LocationDto, error) {
	dsnap, err := config.GetFirestore().Collection("locations").Doc(ID).Get(config.Ctx)
	if err != nil {
		log.Printf("error find location by id: %v\n", err)
	}

	location := domain.Location{}
	err = mapstructure.Decode(dsnap.Data(), &location)
	if err != nil {
		log.Printf("error find location by id: %v\n", err)
	}

	return domain.LocationDto{ID: ID, Location: location}, err
}

func FindAllLocationsByType(LocationType int) ([]domain.LocationDto, error) {
	locationDtos := []domain.LocationDto{}
	iter := config.GetFirestore().Collection("locations").Where("LocationType", "==", LocationType).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find all locations by type: %v\n", err)
			return locationDtos, err
		}

		location := domain.Location{}
		err = mapstructure.Decode(doc.Data(), &location)
		if err != nil {
			log.Printf("error find all locations by type: %v\n", err)
			return locationDtos, err
		}

		locationDtos = append(locationDtos, domain.LocationDto{ID: doc.Ref.ID, Location: location})
	}

	return locationDtos, nil
}

func SaveLocation(location domain.Location) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	ref, wr, err := config.GetFirestore().Collection("locations").Add(config.Ctx, location)
	if err != nil {
		log.Printf("error save location: %v\n", err)
	}

	return ref, wr, err
}
