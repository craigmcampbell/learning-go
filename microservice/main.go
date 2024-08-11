package main

import (
	"os"

	"cc.tech/router"
)

func main() {
	router := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
