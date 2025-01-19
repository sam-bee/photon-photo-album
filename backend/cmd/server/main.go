package main

import (
	"log"
	"os"
	"photo-classifier/internal/api"
	"photo-classifier/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/photos.db"
	}

	db, err := database.NewManager(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize API handlers
	handler := api.NewHandler(db)

	// Set up router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API routes
	v1 := r.Group("/api")
	{
		v1.GET("/photos", handler.GetPhotos)
		v1.GET("/albums", handler.GetAlbums)
		v1.POST("/albums", handler.CreateAlbum)
		v1.PATCH("/albums/:id", handler.UpdateAlbum)
		v1.DELETE("/albums/:id", handler.DeleteAlbum)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
