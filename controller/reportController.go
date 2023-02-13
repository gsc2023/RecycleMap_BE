package controller

import (
	"domain"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"

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
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatalf("[controller:report] error addReport : %v\n", err)
	}

	var data map[string]interface{}
	json.Unmarshal([]byte(value), &data)

	report := domain.Report{}
	err = mapstructure.Decode(data, &report)
	if err != nil {
		log.Fatalf("[controller:report] error addReport : %v\n", err)
	}

	service.Join(report)
}

func toggleLikeOfReport(c *gin.Context) {

}

func delReport(c *gin.Context) {

}

func modifyReport(c *gin.Context) {

}
