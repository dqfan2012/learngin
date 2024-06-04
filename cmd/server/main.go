package main

import (
	"github.com/dqfan2012/learngin/internal/routes"
	"github.com/dqfan2012/learngin/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// Initialize the database
	db.Init()

	// Initialize the Gin router
	r := gin.Default()
	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
