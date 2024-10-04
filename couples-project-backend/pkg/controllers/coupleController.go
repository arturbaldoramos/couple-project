package controllers

import (
	"couples-project-backend/pkg/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCouple(c *gin.Context) {
	var input struct {
		User1UUID string    `json:"user1_uuid" binding:"required"`
		User2UUID string    `json:"user2_uuid" binding:"required"`
		StartDate time.Time `json:"start_date"`
	}

	// Bind JSON to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Criar o objeto Couple
	couple := models.Couple{
		User1UUID: input.User1UUID,
		User2UUID: input.User2UUID,
		StartDate: input.StartDate,
	}

	// Criar o casal no banco de dados
	if coupleResponse, err := couple.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, coupleResponse)
		return
	}
}
