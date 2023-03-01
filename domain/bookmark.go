package domain

type LocationContainer struct {
	LocationID string
}

type BookmarkDao struct {
	UID        string
	LocationID string
	Status     bool
}

type BookmarkDto struct {
	ID       string
	Bookmark BookmarkDao
}
