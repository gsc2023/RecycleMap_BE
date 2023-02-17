package controller

import (
	"github.com/gin-gonic/gin"
)

func authRouter(auth *gin.RouterGroup) {
	auth.POST("/signup", signup)
	auth.POST("/login", signin)
	auth.POST("/email/duplicate", checkEmailDuplicate)
	auth.POST("/nickname/duplicate", checkNickNameDuplicate)
}

func signup(c *gin.Context) {

}

func signin(c *gin.Context) {

}

func checkEmailDuplicate(c *gin.Context) {

}

func checkNickNameDuplicate(c *gin.Context) {

}
