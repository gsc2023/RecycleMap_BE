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
	report.GET("/:reportId", getReport)
	report.POST("/new", addReport)
	report.POST("/:locationId/like", toggleLikeOfReport)
	report.DELETE("/:reportId", delReport)
	report.PATCH("/:reportId", modifyReport)
}

func getAllReport(c *gin.Context) {
	c.JSON(http.StatusOK, service.FindReports())
}

func getReport(c *gin.Context) {

}

func addReport(c *gin.Context) {

	report := domain.Report{}
	err := c.Bind(&report)

	if err != nil {
		log.Fatalf("[controller:report] error addReport : %v\n", err)
	}

	ref, _ := service.Join(report)

	c.String(http.StatusOK, ref.ID)
}

func toggleLikeOfReport(c *gin.Context) {

}

func delReport(c *gin.Context) {

}

func modifyReport(c *gin.Context) {

}
