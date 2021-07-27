package Models

import "gorm.io/gorm"

//Product details
type Product struct {
	Id           int   `json:"id"`
	ProductName string `json:"product_name"`
	Price        uint   `json:"price"`
	Quantity     uint `json:"quantity"`
}
func (b *Product) TableName() string {
	return "products"
}

//Order Details
type Order struct {
	gorm.Model
	Id			int	 `json:"id"`
	CustomerId	int	 `json:"customer_id"`
	ProductId	int	 `json:"product_id"`
	Quantity	uint	 `json:"quantity"`
	Status		string	 `json:"status"`
}

func (b *Order) TableName() string {
return "orders"
}

//Customer Details
type Customer struct{
	CustomerId 	int 	 `json:"customer_id"`
	Name string  `json:"customer_name"`
	Email string `json:"customer_email"`
}

func (b *Customer) TableName() string {
	return "customer"
}

