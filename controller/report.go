package controller

import (
	"domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"service"
)

func reportRouter(report *gin.RouterGroup) {
	report.GET("/", getAllReport)
	report.GET("/:ID", getReport)
	report.POST("/new", addReport)
	report.POST("/:ID/like", toggleLikeOfReport)
	report.DELETE("/:ID", delReport)
	report.PATCH("/:ID", modifyReport)
}

func getAllReport(c *gin.Context) {
	reportDto, err := service.FindReports()

	if err != nil {
		log.Printf("[controller:report] error get all Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, reportDto)
}

func getReport(c *gin.Context) {
	ID := domain.UriParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:report] error get Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	reportDto, err := service.FindReport(ID.ID)

	if err != nil {
		log.Printf("[controller:report] error get Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, reportDto)
}

func addReport(c *gin.Context) {
	tokenString := c.Request.Header.Get("AccessToken")

	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:report] error addReport : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	report := domain.ReportDao{}
	err = c.Bind(&report)

	if err != nil {
		log.Printf("[controller:report] error addReport : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	_, _, err = service.JoinReport(token, report)

	if err != nil {
		log.Printf("[controller:report] error addReport : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}

func toggleLikeOfReport(c *gin.Context) {
	ID := domain.UriParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:report] error toggle like : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	tokenString := c.Request.Header.Get("AccessToken")

	token, err := service.VerifyToken(domain.AccessTokenContainer{AccessToken: tokenString})

	if err != nil {
		log.Printf("[controller:report] error toggle like : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	status, err := service.ToggleLikeOfReport(token, ID.ID)

	if err != nil {
		log.Printf("[controller:report] error toggle like : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, domain.StatusContainer{Status: status})
}

func delReport(c *gin.Context) {
	ID := domain.UriParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:report] error delete Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	_, err = service.DelReport(ID.ID)

	if err != nil {
		log.Printf("[controller:report] error delete Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}

func modifyReport(c *gin.Context) {
	ID := domain.UriParameter{}
	report := domain.ReportDao{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:report] error modify Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	err = c.Bind(&report)

	if err != nil {
		log.Printf("[controller:report] error modify Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	_, err = service.ModifyReport(ID.ID, report)

	if err != nil {
		log.Printf("[controller:report] error modify Report : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}
