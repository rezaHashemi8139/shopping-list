package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	handler := NewHandler(db)
	
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", handler.Register)
	}
}

