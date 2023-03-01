package repository

import (
	"domain"
	"log"
	"module/config"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

func SaveBookmark(bookmark domain.BookmarkDao) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	ref, wr, err := config.GetFirestore().Collection("bookmarks").Add(config.Ctx, bookmark)
	if err != nil {
		log.Printf("error save bookmark: %v\n", err)
	}

	return ref, wr, err
}

func FindValidBookmarks() ([]domain.BookmarkDto, error) {
	bookmarkDtos := []domain.BookmarkDto{}
	iter := config.GetFirestore().Collection("bookmarks").Where("Status", "==", true).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find bookmarks by UID: %v\n", err)
			return bookmarkDtos, err
		}

		bookmark := domain.BookmarkDao{}
		err = mapstructure.Decode(doc.Data(), &bookmark)
		if err != nil {
			log.Printf("error find bookmarks by UID: %v\n", err)
			return bookmarkDtos, err
		}

		bookmarkDtos = append(bookmarkDtos, domain.BookmarkDto{ID: doc.Ref.ID, Bookmark: bookmark})
	}

	return bookmarkDtos, nil
}

func FindBookmarkByID(ID string) (domain.BookmarkDto, error) {
	bookmark := domain.BookmarkDao{}

	dsnap, err := config.GetFirestore().Collection("bookmarks").Doc(ID).Get(config.Ctx)
	if err != nil {
		log.Printf("error find bookmark by id: %v\n", err)
		return domain.BookmarkDto{ID: ID, Bookmark: bookmark}, err
	}

	err = mapstructure.Decode(dsnap.Data(), &bookmark)
	if err != nil {
		log.Printf("error find bookmark by id: %v\n", err)
		return domain.BookmarkDto{ID: ID, Bookmark: bookmark}, err
	}

	return domain.BookmarkDto{ID: ID, Bookmark: bookmark}, err
}

func FindBookmarksByUID(UID string) ([]domain.BookmarkDto, error) {
	bookmarkDtos := []domain.BookmarkDto{}
	iter := config.GetFirestore().Collection("bookmarks").Where("UID", "==", UID).Where("Status", "==", true).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find bookmarks by UID: %v\n", err)
			return bookmarkDtos, err
		}

		bookmark := domain.BookmarkDao{}
		err = mapstructure.Decode(doc.Data(), &bookmark)
		if err != nil {
			log.Printf("error find bookmarks by UID: %v\n", err)
			return bookmarkDtos, err
		}

		bookmarkDtos = append(bookmarkDtos, domain.BookmarkDto{ID: doc.Ref.ID, Bookmark: bookmark})
	}

	return bookmarkDtos, nil
}

func FindBookmarksByLocationID(LocationID string) ([]domain.BookmarkDto, error) {
	bookmarkDtos := []domain.BookmarkDto{}
	iter := config.GetFirestore().Collection("bookmarks").Where("LocationID", "==", LocationID).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find bookmarks by LocationID: %v\n", err)
			return bookmarkDtos, err
		}

		bookmark := domain.BookmarkDao{}
		err = mapstructure.Decode(doc.Data(), &bookmark)
		if err != nil {
			log.Printf("error find bookmarks by LocationID: %v\n", err)
			return bookmarkDtos, err
		}

		bookmarkDtos = append(bookmarkDtos, domain.BookmarkDto{ID: doc.Ref.ID, Bookmark: bookmark})
	}

	return bookmarkDtos, nil
}

func FindBookmarkByUIDAndLocationID(UID string, LocationID string) ([]domain.BookmarkDto, error) {
	bookmarkDtos := []domain.BookmarkDto{}
	iter := config.GetFirestore().Collection("bookmarks").Where("UID", "==", UID).Where("LocationID", "==", LocationID).Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find bookmarks by UID, LocationID: %v\n", err)
			return bookmarkDtos, err
		}

		bookmark := domain.BookmarkDao{}
		err = mapstructure.Decode(doc.Data(), &bookmark)
		if err != nil {
			log.Printf("error find bookmarks by UID, LocationID: %v\n", err)
			return bookmarkDtos, err
		}

		bookmarkDtos = append(bookmarkDtos, domain.BookmarkDto{ID: doc.Ref.ID, Bookmark: bookmark})
	}

	return bookmarkDtos, nil
}

func FindBookmarks() ([]domain.BookmarkDto, error) {
	bookmarkDtos := []domain.BookmarkDto{}
	iter := config.GetFirestore().Collection("bookmarks").Documents(config.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("error find all bookmarks: %v\n", err)
			return bookmarkDtos, err
		}

		bookmark := domain.BookmarkDao{}
		err = mapstructure.Decode(doc.Data(), &bookmark)
		if err != nil {
			log.Printf("error find all bookmarks: %v\n", err)
			return bookmarkDtos, err
		}

		bookmarkDtos = append(bookmarkDtos, domain.BookmarkDto{ID: doc.Ref.ID, Bookmark: bookmark})
	}

	return bookmarkDtos, nil
}

func SetBookmark(ID string, bookmark domain.BookmarkDao) (*firestore.WriteResult, error) {
	wr, err := config.GetFirestore().Collection("bookmarks").Doc(ID).Set(config.Ctx, bookmark)
	if err != nil {
		log.Printf("error set bookmark: %v\n", err)
	}

	return wr, err
}

func DelBookmark(ID string) (*firestore.WriteResult, error) {
	wr, err := config.GetFirestore().Collection("bookmarks").Doc(ID).Delete(config.Ctx)
	if err != nil {
		log.Printf("error delete bookmark: %v\n", err)
	}

	return wr, err
}

func IsBookmarkOwner(UID string, ID string) (bool, error) {
	bookmarkDto, err := FindBookmarkByID(ID)

	if err != nil {
		log.Printf("error owner bookmark: %v\n", err)
		return false, err
	}

	return bookmarkDto.Bookmark.UID == UID, err
}
