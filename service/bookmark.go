package service

import (
	"domain"
	"log"
	"repository"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
)

func FindBookmarks() ([]domain.BookmarkDto, error) {
	return repository.FindValidBookmarks()
}

func FindBookmark(ID string) (domain.BookmarkDto, error) {
	return repository.FindBookmarkByID(ID)
}

func FindMyBookmark(token *auth.Token) ([]domain.BookmarkDto, error) {
	return repository.FindBookmarksByUID(token.UID)
}

func JoinBookmark(token *auth.Token, bookmark domain.BookmarkDao) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	bookmark.UID = token.UID
	return repository.SaveBookmark(bookmark)
}

func DelBookmark(token *auth.Token, ID string) (*firestore.WriteResult, error) {
	err := IsOwner(repository.IsBookmarkOwner(token.UID, ID))

	if err != nil {
		log.Printf("error delete bookmark: %v\n", err)
		return nil, err
	}

	return repository.DelBookmark(ID)
}

func ToggleBookmark(token *auth.Token, ID string) (status bool, err error) {
	bookmarks, err := repository.FindBookmarkByUIDAndLocationID(token.UID, ID)

	if err != nil {
		log.Printf("error toggle bookmark: %v\n", err)
		return false, err
	}

	var bookmarkID string
	toggleBookmark := true

	if len(bookmarks) != 0 {
		toggleBookmark = !bookmarks[0].Bookmark.Status
		bookmarkID = bookmarks[0].ID
	} else { // 저장 안되어있음
		ref, _, err := repository.SaveBookmark(domain.BookmarkDao{})

		if err != nil {
			log.Printf("error toggle bookmark: %v\n", err)
			return false, err
		}

		bookmarkID = ref.ID
	}

	bookmark := domain.BookmarkDao{UID: token.UID, LocationID: ID, Status: toggleBookmark}

	_, err = repository.SetBookmark(bookmarkID, bookmark)

	if err != nil {
		log.Printf("error toggle bookmark: %v\n", err)
		return false, err
	}

	return toggleBookmark, err
}
