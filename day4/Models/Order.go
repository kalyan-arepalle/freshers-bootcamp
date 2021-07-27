//Models/User.go
package Models

import (
	"day4/Config"
)

func CreateOrder(order *Order) error {
	//err := Config.DB.Table("products").Where("id = ? AND quantity >= ?", order.ProductId,order.Quantity).UpdateColumn("quantity",gorm.Expr("quantity - ?",order.Quantity)).Error
	//if err == nil{
	//	return err
	//}
	if err1 := Config.DB.Table("orders").Create(order).Error; err1 != nil{
		return err1
	}
	return nil
}

func GetOrderByID(order *Order, id string) error{
	if err := Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Table("orders").Find(order).Error; err != nil {
		return err
	}
	return nil
}