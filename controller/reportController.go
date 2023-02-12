package controller

import (
	"service"

	"github.com/gin-gonic/gin"
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
	c.String(200, "OK")

	service.Example()
}

func getReport(c *gin.Context) {

}

func addReport(c *gin.Context) {

}

func toggleLikeOfReport(c *gin.Context) {

}

func delReport(c *gin.Context) {

}

func modifyReport(c *gin.Context) {

}
