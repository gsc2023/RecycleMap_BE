package domain

type LikeDao struct {
	UID      string
	ReportID string
	Status   bool
}

type LikeDto struct {
	ID   string
	Like LikeDao
}
