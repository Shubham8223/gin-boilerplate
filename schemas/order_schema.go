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

type ResponseOrderOutput struct {
	User      string  `json:"user"`
	Product   string  `json:"product"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}
type ResponseOrderListOutput struct {
	Product   string  `json:"product"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}
type ResponseOrderPerUserOutput struct {
	User      string  `json:"user"`
	Orders    []ResponseOrderListOutput `json:"orders"`
}