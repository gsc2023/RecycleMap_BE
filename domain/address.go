package domain

type AddressRequest struct {
	Latitude  float64
	Longitude float64
}

type AddressResult struct {
	Text string `json:"text"`
}

type AddressService struct {
	Result []AddressResult `json:"result"`
}

type AddressResponse struct {
	Service AddressService `json:"service"`
}

type AddressReturn struct {
	Address interface{}
}
