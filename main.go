package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rezaHashemi8139/shopping-list/config"
	"github.com/rezaHashemi8139/shopping-list/db"
	"github.com/rezaHashemi8139/shopping-list/modules/user"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it, using default/environment values")
	}

	cfg := config.Load()
	db.Connect(cfg)

	if err := db.DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Register user routes
	user.RegisterRoutes(r, db.DB)

	// Handle 404 - route not found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error":   "Not Found",
			"message": "متأسفانه صفحه مورد نظر یافت نشد!",
			"path":    c.Request.URL.Path,
		})
	})

	r.Run(":" + cfg.Port)
}
