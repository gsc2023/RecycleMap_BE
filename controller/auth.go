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
	//auth.POST("/login", signin)
	auth.POST("/email/duplicate", checkEmailDuplicate)
	auth.POST("/nickname/duplicate", checkNickNameDuplicate)
}

func signup(c *gin.Context) {
	user := domain.User{}

	if err := c.Bind(&user); err != nil {
		log.Printf("[controller:user] error signup : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	userRecord, err := service.JoinUser(user)

	if err != nil {
		log.Printf("[controller:user] error signup : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, userRecord)
}

func signin(c *gin.Context) {
	signinRequestDto := domain.SigninRequestDto{}

	if err := c.Bind(&signinRequestDto); err != nil {
		log.Printf("[controller:user] error signin : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	token, err := service.SignIn(signinRequestDto)

	if err != nil {
		log.Printf("[controller:user] error signin : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	response := domain.AccessTokenContainer{
		AccessToken: token,
	}

	c.JSON(http.StatusOK, response)
}

func checkEmailDuplicate(c *gin.Context) {

}

func checkNickNameDuplicate(c *gin.Context) {

}
