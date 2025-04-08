package services

import (
	"errors"

	"gin-boilerplate/config"
	"gin-boilerplate/models"
	"gin-boilerplate/schemas"
	"gin-boilerplate/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(tx *gorm.DB, input *schemas.CreateUserInput) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil,err
	}
	input.Password = string(hashedPassword)
	user := &models.User{}
	utils.StructToMapCreate(input, user)

	if err := tx.Create(user).Error; err != nil {
		return nil,err
	}

	return user, nil
}

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := config.DB.Preload("Orders.Product").Preload("Orders").First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}


func UpdateUser(userID uint, updatedData *schemas.UpdateUserInput) error {
	updates := utils.StructToMapUpdate(updatedData)

	if len(updates) == 0 {
		return nil 
	}
	return config.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error
}

func DeleteUser(userID uint) error {
	return config.DB.Delete(&models.User{}, userID).Error
}