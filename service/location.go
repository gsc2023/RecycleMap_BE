package service

import (
	"domain"
	"repository"

	"cloud.google.com/go/firestore"
)

func FindLocations() ([]domain.LocationDto, error) {
	return repository.FindAllLocations()
}

func FindLocation(ID string) (domain.LocationDto, error) {
	return repository.FindLocationById(ID)
}

func SaveLocation(location domain.Location) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return repository.SaveLocation(location)
}
