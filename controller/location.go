package controller

import (
	"domain"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func locationRouter(location *gin.RouterGroup) {
	location.GET("/", getAll)
	location.GET("/:locationId", findById)
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

func setBookmark(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getCommentsByLocationId(c *gin.Context) {
	c.Status(http.StatusOK)
}

func saveCommentToLocation(c *gin.Context) {
	c.Status(http.StatusOK)
}

func updateComment(c *gin.Context) {
	c.Status(http.StatusOK)
}

func deleteComment(c *gin.Context) {
	c.Status(http.StatusOK)
}
