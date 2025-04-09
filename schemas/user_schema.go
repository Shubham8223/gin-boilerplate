package schemas

import "github.com/lib/pq"

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=admin user queen"` 
}

type UpdateUserInput struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Role     *string `json:"role"`
}
type ResponseUserOutput struct {
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	OrderIds pq.Int64Array  `json:"order_ids" gorm:"type:integer[]"` 
}
