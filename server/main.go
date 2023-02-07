package main

import (
	"context"

	"module/config"

	"log"
)

func main() {
	log.Println(1111)
	app := config.InitializeAppWithServiceAccount()

	log.Println(app)

	ctx := context.Background()

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
