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

func LoginUser(authUser *schemas.AuthUser) (string, error) {
	var user models.User

	if err := config.DB.Where("email = ?", authUser.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("user not found")
		}
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authUser.Password)); err != nil {
       return "", err
    }

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
