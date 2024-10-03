package main

import (
	database "couples-project-backend/pkg/db"
	"couples-project-backend/pkg/models"
	"couples-project-backend/pkg/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Connect to database
	database.Connect()

	// Migrate database
	migratironErr := database.DB.AutoMigrate(&models.User{}, &models.Couple{})
	if migratironErr != nil {
		panic("Error migrating database")
	}

	// Initialize routes
	r := gin.Default()
	r.Use(gin.Logger())
	router.InitializeRoutes(r)

	r.Run()
}
