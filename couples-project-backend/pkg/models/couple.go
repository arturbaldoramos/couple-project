package models

import (
	"time"

	"gorm.io/gorm"
)

type Couple struct {
	gorm.Model

	ID        int       `gorm:"primary_key;type:integer;not null" json:"id"`
	User1     User      `gorm:"foreignKey:User1UUID" json:"user1"`
	User2     User      `gorm:"foreignKey:User2UUID" json:"user2"`
	User1UUID string    `gorm:"type:uuid;not null" json:"user1_uuid"`
	User2UUID string    `gorm:"type:uuid;not null" json:"user2_uuid"`
	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
