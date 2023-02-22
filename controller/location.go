package controller

import (
	"domain"
	"log"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func locationRouter(location *gin.RouterGroup) {
	location.GET("/", getAllLocation)
	location.GET("/:locationId", getLocation)
	location.POST("/new", saveLocation)
	location.POST("/:locationId/bookmark", setBookmark)
	location.GET("/:locationId/comments", getCommentsByLocationId)
	location.POST("/:locationId/comments", saveCommentToLocation)
	location.PATCH("/comments/:commentId", updateComment)
	location.DELETE("/comments/:commentId", deleteComment)
}

func getAllLocation(c *gin.Context) {
	locationDto, err := service.FindLocations()

	if err != nil {
		log.Printf("[controller:location] error getAllLocation : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, locationDto)

}

func getLocation(c *gin.Context) {
	ID := domain.LocationUrlParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:location] error getLocation : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	locationDto, err := service.FindLocation(ID.ID)

	if err != nil {
		log.Printf("[controller:report] error getLocation : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, locationDto)
}

func saveLocation(c *gin.Context) {
	location := domain.Location{}
	err := c.Bind(&location)

	if err != nil {
		log.Printf("[controller:location] error saveLocation : %v\n", err)
	}

	ref, _, err := service.SaveLocation(location)

	if err != nil {
		log.Printf("[controller:location] error saveLocation : %v\n", err)
	}

	c.JSON(http.StatusOK, ref.ID)
}

func setBookmark(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getCommentsByLocationId(c *gin.Context) {
	c.Status(http.StatusOK)
}

func saveCommentToLocation(c *gin.Context) {
	ID := domain.LocationUrlParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:location] error saveCommentToLocation : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	tokenString := c.Request.Header["AccessToken"]
	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString[0]})

	if err != nil {
		log.Printf("[controller:location] error saveCommentToLocation : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	comment := domain.Comment{}
	err = c.Bind(&comment)

	if err != nil {
		log.Printf("[controller:location] error saveCommentToLocation : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	_, _, err = service.JoinComment(token, ID.ID, comment)

	if err != nil {
		log.Printf("[controller:location] error saveCommentToLocation : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}

func updateComment(c *gin.Context) {
	c.Status(http.StatusOK)
}

func deleteComment(c *gin.Context) {
	c.Status(http.StatusOK)
}
