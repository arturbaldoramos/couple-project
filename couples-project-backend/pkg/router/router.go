package router

import (
	"couples-project-backend/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/user/:uuid", controllers.GetUserByUUID)
	router.POST("/user", controllers.CreateUser)

	router.POST("/couple", controllers.CreateCouple)
}
