package router

import (
	"couples-project-backend/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser)
}
