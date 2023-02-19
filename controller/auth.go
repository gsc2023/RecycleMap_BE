package controller

import (
	"domain"
	"log"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func authRouter(auth *gin.RouterGroup) {
	auth.POST("/signup", signup)
	auth.POST("/login", signin)
	auth.POST("/email/duplicate", checkEmailDuplicate)
	auth.POST("/nickname/duplicate", checkNickNameDuplicate)
}

func signup(c *gin.Context) {
	user := domain.User{}
	err := c.Bind(&user)

	if err != nil {
		log.Printf("[controller:user] error signup : %v\n", err)
	}

	c.JSON(http.StatusOK, service.JoinUser(user))
}

func signin(c *gin.Context) {

}

func checkEmailDuplicate(c *gin.Context) {

}

func checkNickNameDuplicate(c *gin.Context) {

}
