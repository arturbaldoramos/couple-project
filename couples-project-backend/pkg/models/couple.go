package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	database "couples-project-backend/pkg/db"
)

type Couple struct {
	gorm.Model

	ID        int       `gorm:"primary_key;type:integer;not null" json:"id"`
	User1     User      `gorm:"foreignKey:User1UUID; references:UUID" json:"user1"`
	User2     User      `gorm:"foreignKey:User2UUID; references:UUID" json:"user2"`
	User1UUID string    `gorm:"type:uuid;not null" json:"user1_uuid"`
	User2UUID string    `gorm:"type:uuid;not null" json:"user2_uuid"`
	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type CoupleResponse struct {
	ID        int       `json:"id"`
	User1UUID string    `json:"user1_uuid"`
	User2UUID string    `json:"user2_uuid"`
	StartDate time.Time `json:"start_date"`
}

func (couple *Couple) Create() (*CoupleResponse, error) {
	// Verificar se o usuário 1 existe
	var user1 User
	if err := database.DB.Where("uuid = ?", couple.User1UUID).First(&user1).Error; err != nil {
		return nil, fmt.Errorf("user1 not found")
	}

	// Verificar se o usuário 2 existe
	var user2 User
	if err := database.DB.Where("uuid = ?", couple.User2UUID).First(&user2).Error; err != nil {
		return nil, fmt.Errorf("user2 not found")
	}

	// Criar a data de início do casal (se não foi passada, usar a data atual)
	if couple.StartDate.IsZero() {
		couple.StartDate = time.Now()
	}

	// Salvar o casal no banco de dados
	if err := database.DB.Create(couple).Error; err != nil {
		return nil, fmt.Errorf("error creating couple: %v", err)
	}

	// Retornar o objeto de resposta
	return &CoupleResponse{
		ID:        couple.ID,
		User1UUID: couple.User1UUID,
		User2UUID: couple.User2UUID,
		StartDate: couple.StartDate,
	}, nil
}
