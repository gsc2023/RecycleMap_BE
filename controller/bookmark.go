package controller

import (
	"domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"service"
)

func bookmarkRouter(bookmark *gin.RouterGroup) {
	bookmark.GET("/", getAllBookmark)
	bookmark.POST("/:ID", toggleBookmark)
}

func getAllBookmark(c *gin.Context) {
	bookmarkDtos, err := service.FindBookmarks()

	if err != nil {
		log.Printf("[controller:bookmark] error get all Bookmark : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, bookmarkDtos)
}

func toggleBookmark(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")

	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:bookmark] error toggle like : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	locationID := domain.LocationContainer{}
	err = c.Bind(&locationID)

	if err != nil {
		log.Printf("[controller:bookmark] error toggle like : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	status, err := service.ToggleBookmark(token, locationID.LocationID)

	if err != nil {
		log.Printf("[controller:bookmark] error toggle like : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, domain.StatusContainer{Status: status})
}
