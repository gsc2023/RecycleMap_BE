package main

import (
	"controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controller.InitRouter(router)

	router.Run(":8080")
}
