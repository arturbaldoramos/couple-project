package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	database "couples-project-backend/pkg/db"
)

type User struct {
	gorm.Model

	ID        uint   `gorm:"primary_key;autoIncrement" json:"id"`
	UUID      string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uuid"`
	Name      string `gorm:"type:varchar(100);not null" validate:"required,min=3,max=100" json:"name"`
	Email     string `gorm:"type:varchar(120);unique;not null;" validate:"required,email,max=120" json:"email"`
	Password  string `gorm:"type:varchar(200);not null" validate:"required,min=8,max=120" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type UserResponse struct {
	UUID      string         `json:"uuid"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (user *User) Create() (*UserResponse, error) {

	//Check if email already exist on database, even if user was deleted
	if err := database.DB.Unscoped().Where("email = ?", user.Email).First(&User{}).Error; err == nil {
		return nil, fmt.Errorf("email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password")
	}
	user.Password = string(hashedPassword)

	// Save the user to the database
	if err := database.DB.Create(user).Error; err != nil {
		return nil, fmt.Errorf("error creating user")
	}

	// Return the user response without the password field and ID
	return &UserResponse{
		UUID:      user.UUID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, nil
}

func (user *User) Read(uuid string) (*UserResponse, error) {
	if err := database.DB.Where("uuid = ?", uuid).First(user).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &UserResponse{
		UUID:      user.UUID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, nil
}
