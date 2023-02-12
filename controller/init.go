package controller

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	reportRouter(router.Group("/report"))

}

func AuthRouter(router *gin.Engine) {
	authRouter(router.Group("/auth"))

}
