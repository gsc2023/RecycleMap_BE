package repository

import (
	"domain"
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

const DISTANCE_FROM_MY_LOCATION = float64(0.0003)

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

func FindAllLocationsByPosition(latitude float64, longitude float64) ([]domain.LocationDto, error) {
	log.Println(latitude, longitude)
	locationDtos := []domain.LocationDto{}

	distance := DISTANCE_FROM_MY_LOCATION
	MaxLat := latitude + distance
	MinLat := latitude - distance
	MaxLong := longitude + distance
	MinLong := longitude - distance
	log.Println(MaxLat, MinLat, MaxLong, MinLong)
	iter := config.GetFirestore().Collection("locations").Where("Latitude", "<=", MaxLat).Where("Latitude", ">=", MinLat).Documents(config.Ctx)
	log.Println(iter)

	iter = config.GetFirestore().Collection("locations").Where("Longitude", ">=", MinLong).Where("Longitude", "<=", MaxLong).Documents(config.Ctx)
	//iter = config.GetFirestore().Collection("locations").Where("Latitude", "<=", MaxLat).Where("Latitude", ">=", MinLat).Where("Longitude", ">=", MinLong).Where("Longitude", "<=", MaxLong).Documents(config.Ctx)
	log.Println(iter)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find all locations by position: %v\n", err)
			return locationDtos, err
		}

		location := domain.Location{}
		err = mapstructure.Decode(doc.Data(), &location)
		if err != nil {
			log.Printf("error find all locations by position: %v\n", err)
			return locationDtos, err
		}

		locationDtos = append(locationDtos, domain.LocationDto{ID: doc.Ref.ID, Location: location})

	}

	iter = config.GetFirestore().Collection("reports").Where("Latitude", "<=", MaxLat).Where("Latitude", ">=", MinLat).Documents(config.Ctx)
	log.Println(iter)

	iter = config.GetFirestore().Collection("reports").Where("Longitude", ">=", MinLong).Where("Longitude", "<=", MaxLong).Documents(config.Ctx)
	log.Println(iter)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find all locations by position: %v\n", err)
			return locationDtos, err
		}

		report := domain.ReportDao{}
		err = mapstructure.Decode(doc.Data(), &report)
		if err != nil {
			log.Printf("error find all locations by position: %v\n", err)
			return locationDtos, err
		}

		locationDtos = append(locationDtos, domain.LocationDto{ID: doc.Ref.ID, Location: domain.Location{
			Name:         report.Name,
			LocationType: report.LocationType,
			Latitude:     report.Latitude,
			Longitude:    report.Longitude,
			Content:      report.Content,
			ImagePath:    report.ImagePath,
		}})

	}

	return locationDtos, nil
}
