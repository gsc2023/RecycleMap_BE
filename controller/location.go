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
	location.GET("/find/:locationType", getAllLocationByType)
	location.POST("/:locationId/bookmark", setBookmark)
	location.GET("/:locationId/comments", getAllCommentByLocationId)
	location.POST("/:locationId/comments", saveCommentToLocation)
	location.PATCH("/comments/:commentId", updateComment)
	location.DELETE("/comments/:commentId", deleteComment)
}

func getAllLocation(c *gin.Context) {
	locationDto, err := service.FindLocations()

	if err != nil {
		log.Printf("[controller:location] error get locations : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, locationDto)

}

func getLocation(c *gin.Context) {
	ID := domain.LocationUrlParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:location] error get location : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	locationDto, err := service.FindLocationById(ID.ID)

	if err != nil {
		log.Printf("[controller:report] error get location : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, locationDto)
}

func getAllLocationByType(c *gin.Context) {
	Type := domain.LocationTypeUrlParameter{}
	err := c.ShouldBindUri(&Type)

	if err != nil {
		log.Printf("[controller:location] error get locations by type : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	locationDto, err := service.FindLocationsByType(Type.LocationType)

	if err != nil {
		log.Printf("[controller:location] error get locations by type : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, locationDto)
}

func saveLocation(c *gin.Context) {
	location := domain.Location{}
	err := c.Bind(&location)

	if err != nil {
		log.Printf("controller:location] error save location : %v\n", err)
	}

	_, _, err = service.SaveLocation(location)

	if err != nil {
		log.Printf("controller:location] error save location : %v\n", err)
	}

	c.Status(http.StatusOK)
}

func setBookmark(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getAllCommentByLocationId(c *gin.Context) {
	ID := domain.LocationUrlParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:location] error get comment by locationId : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	commentDto, err := service.FindCommentsById(ID.ID)

	if err != nil {
		log.Printf("[controller:location] error get comment by locationId : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, commentDto)
}

func saveCommentToLocation(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")
	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:location] error save comment to location : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	ID := domain.LocationUrlParameter{}
	err = c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:location] error save comment to location : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	comment := domain.Comment{}
	err = c.Bind(&comment)

	if err != nil {
		log.Printf("[controller:location] error save comment to location : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	_, _, err = service.JoinComment(token, ID.ID, comment)

	if err != nil {
		log.Printf("[controller:location] error save comment to location : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}

// func updateComment(c *gin.Context) {
// 	c.Status(http.StatusOK)
// }

// func deleteComment(c *gin.Context) {
// 	c.Status(http.StatusOK)
// }
