package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UUID      string `gorm:"primary_key; type:uuid;default:gen_random_uuid()" json:"uuid"`
	ID        int    `gorm:"type:integer;not null" json:"id"`
	Name      string `gorm:"type:varchar(100);not null" validate:"required,min=3,max=100" json:"name"`
	Email     string `gorm:"type:varchar(120);unique;not null;" validate:"required,email,max=120" json:"email"`
	Password  string `gorm:"type:varchar(200);not null" validate:"required,min=8,max=120" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (user *User) Create() {

}
