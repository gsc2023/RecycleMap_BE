package service

import (
	"domain"
	"repository"

	"cloud.google.com/go/firestore"
)

func FindLocations() ([]domain.LocationDto, error) {
	return repository.FindAllLocations()
}

func FindLocationById(ID string) (domain.LocationDto, error) {
	return repository.FindLocationById(ID)
}
func FindLocationsByType(LocationType int) ([]domain.LocationDto, error) {
	return repository.FindAllLocationsByType(LocationType)
}

func SaveLocation(location domain.Location) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return repository.SaveLocation(location)
}

func FindAroundLocations(location domain.MyLocation) ([]domain.LocationDto, error) {
	return repository.FindAllLocationsByPosition(location.Latitude, location.Longitude)
}
