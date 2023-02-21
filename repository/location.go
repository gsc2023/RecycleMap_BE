package repository

import (
	"domain"
	"log"
	"module/config"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

func FindAllLocations() (locationDtos []domain.LocationDto) {
	iter := config.GetFirestore().Collection("locations").Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find all locations: %v\n", err)
		}

		location := domain.Location{}
		err = mapstructure.Decode(doc.Data(), &location)
		if err != nil {
			log.Printf("error find all locations: %v\n", err)
		}

		locationDtos = append(locationDtos, domain.LocationDto{ID: doc.Ref.ID, Location: location})
	}

	return
}

func FindLocationByID(ID string) domain.LocationDto {
	dsnap, err := config.GetFirestore().Collection("locations").Doc(ID).Get(config.Ctx)
	if err != nil {
		log.Printf("error find location by id: %v\n", err)
	}

	location := domain.Location{}
	err = mapstructure.Decode(dsnap.Data(), &location)
	if err != nil {
		log.Printf("error find location by id: %v\n", err)
	}

	return domain.LocationDto{ID: ID, Location: location}
}
