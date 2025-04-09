package enums

import "gin-boilerplate/models"

type ObacResource struct {
	Model     interface{}
	OwnerKey  string
	TableName string
}

var ResourceMap = map[string]ObacResource{
	"orders": {
		Model:    &models.Order{},
		OwnerKey: "UserID",
		TableName: "orders",
	},
	"users": {
		Model:    &models.User{},
		OwnerKey: "ID",
		TableName: "users",
	},
}
