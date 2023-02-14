package controller

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	reportRouter(router.Group("/report"))
	authRouter(router.Group("/auth"))

}
