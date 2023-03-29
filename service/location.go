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

func FindAroundLocations(location domain.MyLocation) ([]domain.LocationDtoWithAddress, error) {
	locationDtos, err := repository.FindAllLocationsByPosition(location.Latitude, location.Longitude)

	if err != nil {
		return nil, err
	}

	locationDtoWithADdresses := []domain.LocationDtoWithAddress{}

	for _, val := range locationDtos {
		address, err := FindAddress(domain.AddressRequest{
			Latitude:  val.Location.Latitude,
			Longitude: val.Location.Longitude,
		})

		if err != nil {
			return nil, err
		}

		locationDtoWithADdresses = append(locationDtoWithADdresses, domain.LocationDtoWithAddress{
			LocationDto: val,
			Address:     address,
		})
	}

	return locationDtoWithADdresses, nil
}
