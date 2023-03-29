package domain

type AddressRequest struct {
	Latitude  float64
	Longitude float64
}

type AddressReturn struct {
	Address interface{}
}
