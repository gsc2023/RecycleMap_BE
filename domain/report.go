package domain

type Report struct {
	UserID         string
	Name           string
	LocationType   int
	Latitude       float64
	Longitude      float64
	Content        string
	Recommendation int
}
