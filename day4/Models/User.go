//Models/User.go
package Models

import (
	"day4/Config"
	"fmt"

)
//GetAllProducts Fetch all products data
func GetAllProducts(prod *[]Product) (err error) {
	if err = Config.DB.Table("products").Find(prod).Error; err != nil {
		return err
	}
	return nil
}
//CreateProduct ... Insert New product
func CreateProduct(prod *Product) (err error) {
	if err = Config.DB.Table("products").Create(prod).Error; err != nil {
		return err
	}
	return nil
}
//GetProductByID ... Fetch only one product by Id
func GetProductByID(prod *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(prod).Error; err != nil {
		return err
	}
	return nil
}
//UpdateProduct ... Update user
func UpdateProduct(prod *Product, id string) (err error) {
	fmt.Println(prod)
	Config.DB.Save(prod)
	return nil
}
//DeleteProduct ... Delete the product
func DeleteProduct(prod *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(prod)
	return nil
}

//CreateCustomer ... Insert New customer
func CreateCustomer(cust *Customer) (err error) {
	if err = Config.DB.Table("customer").Create(cust).Error; err != nil {
		return err
	}
	return nil
}
func GetCustomerByID(customer *Customer, id string) (err error) {
	if err = Config.DB.Where("customer_id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}