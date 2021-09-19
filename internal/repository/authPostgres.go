package repository

import (
	"github.com/hostingvk4/badgerList/internal/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (uint, error) {
	result := r.db.Create(&user)

	return uint(user.ID), result.Error
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var userModel models.User
	err := r.db.Where("username = ? AND password = ?", username, password).First(&userModel).Error

	return userModel, err
}

func (r *AuthPostgres) SetRefreshToken(userId uint, token models.RefreshToken) error {
	result := r.db.Create(&token)

	return result.Error
}
