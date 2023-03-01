package controller

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	reportRouter(router.Group("/reports"))
	authRouter(router.Group("/auth"))
	locationRouter(router.Group("/locations"))
	userRouter(router.Group("/my"))
	bookmarkRouter(router.Group("/bookmarks"))
}
