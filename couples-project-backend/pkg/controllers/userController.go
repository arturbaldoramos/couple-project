package controllers

import (
	"couples-project-backend/pkg/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetUserByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	user := new(models.User)

	// Busca o usuário pelo UUID
	userResponse, err := user.Read(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Retorna a resposta do usuário sem a senha
	c.JSON(http.StatusOK, userResponse)
}

func CreateUser(c *gin.Context) {
	user := new(models.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	if userResponse, err := user.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, userResponse)
		return
	}
}
