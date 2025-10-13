package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rezaHashemi8139/shopping-list/config"
	"github.com/rezaHashemi8139/shopping-list/db"
	"github.com/rezaHashemi8139/shopping-list/modules/user"
)

func main() {
	cfg := config.Load()
	db.Connect(cfg)

	// مهاجرت خودکار جداول اولیه
	if err := db.DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	// راه‌اندازی Gin
	r := gin.Default()

	// یک روت ساده برای تست
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// TODO: register module routes (بعداً)
	// routes.Register(r, cfg)

	r.Run(":" + cfg.Port)
}
