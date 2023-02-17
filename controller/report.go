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
	c.JSON(http.StatusOK, service.FindReports())
}

func getReport(c *gin.Context) {
	ID := domain.UriParameter{}
	c.ShouldBindUri(&ID)

	c.JSON(http.StatusOK, service.FindReport(ID.ID))
}

func addReport(c *gin.Context) {
	report := domain.Report{}
	err := c.Bind(&report)

	if err != nil {
		log.Printf("[controller:report] error addReport : %v\n", err)
	}

	ref, _ := service.JoinReport(report)

	c.String(http.StatusOK, ref.ID)
}

func toggleLikeOfReport(c *gin.Context) {

}

func delReport(c *gin.Context) {
	ID := domain.UriParameter{}
	c.ShouldBindUri(&ID)

	service.DelReport(ID.ID)

	c.Status(http.StatusOK)
}

func modifyReport(c *gin.Context) {
	ID := domain.UriParameter{}
	report := domain.Report{}
	c.ShouldBindUri(&ID)
	err := c.Bind(&report)

	if err != nil {
		log.Printf("[controller:report] error modifyReport : %v\n", err)
	}

	service.ModifyReport(ID.ID, report)

	c.Status(http.StatusOK)
}
