package main

import (
	"log"

	"module/config"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})
	return r
}

func main() {
	// Use a service account
	app := config.InitializeAppWithServiceAccount()

	client, err := app.Firestore(config.Ctx)
	if err != nil {
		log.Println("firestore")
		log.Fatalln(err)
	}

	_, _, err = client.Collection("users").Add(config.Ctx, map[string]interface{}{
		"first": "alallala",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	defer client.Close()

	r := setupRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
