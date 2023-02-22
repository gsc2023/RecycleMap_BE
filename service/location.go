package service

import (
	"domain"
	"repository"

	"cloud.google.com/go/firestore"
)

func FindLocations() (rocationDtos []domain.LocationDto) {
	return repository.Find()
}

func FindLocation(ID string) domain.LocationDto {
	return repository.FindOne(ID)
}

func SaveLocation(location domain.Location) (*firestore.DocumentRef, *firestore.WriteResult) {
	return repository.Save(location)
}
