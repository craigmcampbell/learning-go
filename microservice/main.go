package main

import (
	"os"
)

func main() {
	router := SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
