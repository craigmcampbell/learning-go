package main

import (
	"log"
	"os"

	"cc.tech/main/config"
	"cc.tech/main/db"
	// "your-project/repository"
)

func main() {
	cfg := config.LoadConfig()
    
	_, err := db.NewDatabase(cfg)
	if err != nil {
			// Handle error
			log.Fatal(err)
	}

	// userRepo := repository.NewUserRepository(database)
	
	// Use userRepo in your application...

	router := SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
