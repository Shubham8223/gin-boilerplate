package schemas

type CreateProductInput struct {
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
}

type UpdateProductInput struct {
	Name       *string  `json:"name"`
	Price      *float64 `json:"price"`
	CategoryID *uint    `json:"category_id"`
}
