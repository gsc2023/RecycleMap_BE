package service

import (
	"domain"
	"repository"
)

func FindLocations() (rocationDtos []domain.LocationDto) {
	return repository.FindAllLocations()
}

func FindLocation(ID string) domain.LocationDto {
	return repository.FindLocationByID(ID)
}
