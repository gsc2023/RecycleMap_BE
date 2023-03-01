package controller

import (
	"domain"
	"log"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func userRouter(user *gin.RouterGroup) {
	user.GET("/comments", getAllMyLocation)
	user.PATCH("/comments/:commentId", updateComment)
	user.DELETE("/comments/:commentId", deleteComment)
	user.GET("/report", getMyReport)
	user.DELETE("/report/:ID", delReport)
	user.PATCH("/report/:ID", modifyReport)
	user.GET("/like", getMyLikePlace)
	user.PATCH("/edit", editUser)
	user.GET("/bookmark", getMyBookmark)
}

func getMyBookmark(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")
	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:user] error get my bookmark : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	bookmarkDtos, err := service.FindMyBookmark(token)

	if err != nil {
		log.Printf("[controller:user] error get my bookmark : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, bookmarkDtos)
}

func getAllMyLocation(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")
	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:user] error get all my location : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	commentDto, err := service.FindCommentsByUID(token)

	if err != nil {
		log.Printf("[controller:user] error get all my location : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, commentDto)
}

func updateComment(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")
	_, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:user] error update my comment : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}

func deleteComment(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")
	_, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:user] error delete my comment : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}

func getMyReport(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")

	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:my] error get my report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	reportDtos, err := service.FindMyReport(token)

	if err != nil {
		log.Printf("[controller:my] error get my report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, reportDtos)
}

func delReport(c *gin.Context) {
	ID := domain.UriParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:my] error delete Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	tokenString := c.Request.Header.Get("AccessToken")

	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:my] error delete Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	_, err = service.DelReport(token, ID.ID)

	if err != nil {
		log.Printf("[controller:my] error delete Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}

func modifyReport(c *gin.Context) {
	ID := domain.UriParameter{}
	report := domain.ReportDao{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:my] error modify Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	err = c.Bind(&report)

	if err != nil {
		log.Printf("[controller:my] error modify Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	tokenString := c.Request.Header.Get("AccessToken")

	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:my] error modify Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	_, err = service.ModifyReport(token, ID.ID, report)

	if err != nil {
		log.Printf("[controller:my] error modify Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}

func getMyLikePlace(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")

	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:my] error modify Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	likeDtos, err := service.FindMyLikeReport(token)

	if err != nil {
		log.Printf("[controller:my] error modify Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	var reportDtos []domain.ReportDto

	for _, val := range likeDtos {
		reportDto, err := service.FindReport(val.Like.ReportID)

		if err != nil {
			log.Printf("[controller:my] error modify Report : %v\n", err)
			c.JSON(http.StatusNotFound, err)
		}

		reportDtos = append(reportDtos, reportDto)
	}

	c.JSON(http.StatusOK, reportDtos)
}

func editUser(c *gin.Context) {
	user := domain.User{}

	if err := c.Bind(&user); err != nil {
		log.Printf("[controller:user] error edit user : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	tokenString := c.Request.Header.Get("AccessToken")

	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:my] error edit user : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	userRecord, err := service.UpdateUser(token, user)

	if err != nil {
		log.Printf("[controller:user] error edit user : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, userRecord)
}
