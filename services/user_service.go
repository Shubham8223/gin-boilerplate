package services

import (
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

func GetUserByID(userID uint) (*schemas.ResponseUserOutput, error) {
	var user schemas.ResponseUserOutput
	query := `
		SELECT u.name, u.email, array_agg(o.id) as order_ids
		FROM users u
		LEFT JOIN orders o ON u.id = o.user_id
		WHERE u.id = ?
		GROUP BY u.id
	`
	if err := config.DB.Raw(query, userID).Scan(&user).Error; err != nil {
		return nil, err
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