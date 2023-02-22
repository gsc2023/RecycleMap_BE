package controller

import (
	"domain"
	"log"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func locationRouter(location *gin.RouterGroup) {
	location.GET("/", getAll)
	location.GET("/:locationId", findById)
	location.POST("/new", saveLocation)
	location.POST("/:locationId/bookmark", setBookmark)
	location.GET("/:locationId/comments", getCommentsByLocationId)
	location.POST("/:locationId/comments", saveCommentToLocation)
	location.PATCH("/comments/:commentId", updateComment)
	location.DELETE("/comments/:commentId", deleteComment)
}

func getAll(c *gin.Context) {
	c.JSON(http.StatusOK, service.FindLocations())
}

func findById(c *gin.Context) {
	ID := domain.LocationUrlParameter{}
	c.ShouldBindUri(&ID)

	c.JSON(http.StatusOK, service.FindLocation(ID.ID))
}
func saveLocation(c *gin.Context) {
	location := domain.Location{}
	err := c.Bind(&location)

	if err != nil {
		log.Printf("[controller:location] error createLocation : %v\n", err)
	}

	ref, _ := service.SaveLocation(location)

	c.String(http.StatusOK, ref.ID)
}
func setBookmark(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getCommentsByLocationId(c *gin.Context) {
	c.Status(http.StatusOK)
}

func saveCommentToLocation(c *gin.Context) {
	locationRef := domain.LocationUrlParameter{}
	c.ShouldBindUri(&locationRef)

	content := domain.Comment{}.Content
	uID := domain.Comment{}.UID

	comment :={locationRef, content, uID}
	err := c.Bind(&content)
	if err != nil {
		log.Printf("[controller:location] error saveCommentToLocation : %v\n", err)
	}
	ref, _ := service.SaveComment(locationRef.ID, content, uID)
	c.String(http.StatusOK, ref.ID)
}

func updateComment(c *gin.Context) {
	c.Status(http.StatusOK)
}

func deleteComment(c *gin.Context) {
	c.Status(http.StatusOK)
}
