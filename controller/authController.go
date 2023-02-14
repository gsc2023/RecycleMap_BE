package controller

import (
	"domain"
	"encoding/json"
	"io/ioutil"
	"log"
	"service"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func authRouter(report *gin.RouterGroup) {
	report.POST("/signup", signup)
	report.POST("/login", login)
	report.POST("/email/duplicate", isEmail)
	report.POST("/nickname/duplicate", isNickname)
}

func signup(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatalf("[controller:auth] error addReport : %v\n", err)
	}

	var data map[string]interface{}
	json.Unmarshal([]byte(value), &data)

	user := domain.User{}
	err = mapstructure.Decode(data, &user)
	if err != nil {
		log.Fatalf("[controller:auth] error signup : %v\n", err)
	}

	service.Signup(user)
}

func login(c *gin.Context) {

}

func isEmail(c *gin.Context) {

}

func isNickname(c *gin.Context) {

}
