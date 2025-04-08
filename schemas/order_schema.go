package schemas

type CreateOrderInput struct {
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	Total     float64 `json:"total"`
}

type UpdateOrderInput struct {
	ProductID *uint    `json:"product_id"`
	Quantity  *int     `json:"quantity"`
}
