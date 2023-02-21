package controller

import (
	"net/http"

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
	c.Status(http.StatusOK)
}

func findById(c *gin.Context) {
	c.Status(http.StatusOK)
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
