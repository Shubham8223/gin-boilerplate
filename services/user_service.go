package services

import (
	"errors"

	"gin-boilerplate/config"
	"gin-boilerplate/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return config.DB.Create(user).Error
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