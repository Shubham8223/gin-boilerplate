package services

import (
	"errors"

	"gin-boilerplate/config"
	"gorm.io/gorm"
	"gin-boilerplate/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(tx *gorm.DB, user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil,err
	}
	user.Password = string(hashedPassword)

	if err := tx.Create(user).Error; err != nil {
		return nil,err
	}

	return user, nil
}

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}


func UpdateUser(userID uint, updatedData map[string]interface{}) error {
	return config.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updatedData).Error
}

func DeleteUser(userID uint) error {
	return config.DB.Delete(&models.User{}, userID).Error
}