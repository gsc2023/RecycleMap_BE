package controller

import (
	"github.com/gin-gonic/gin"
)

func authRouter(report *gin.RouterGroup) {
	report.POST("/signup", signup)
	report.POST("/login", login)
	report.POST("/email/duplicate", isEmail)
	report.POST("/nickname/duplicate", isNickname)
}

func signup(c *gin.Context) {

}

func login(c *gin.Context) {

}

func isEmail(c *gin.Context) {

}

func isNickname(c *gin.Context) {

}
