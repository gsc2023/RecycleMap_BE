package domain

type ReportDao struct {
	UID          string
	Name         string
	LocationType int
	Latitude     float64
	Longitude    float64
	Content      string
	Like         int
}

type ReportDto struct {
	ID     string
	Report ReportDao
}
