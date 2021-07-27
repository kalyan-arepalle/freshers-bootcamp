//Models/User.go
package Models

import (
	"day4/Config"
)

// CreateOrder Place an order
func CreateOrder(order *Order) error {
	if err1 := Config.DB.Table("orders").Create(order).Error; err1 != nil{
		return err1
	}
	return nil
}

// GetOrderByID Get order by their ID
func GetOrderByID(order *Order, id string) error{
	if err := Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

// GetAllOrders Get all orders
func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Table("orders").Find(order).Error; err != nil {
		return err
	}
	return nil
}