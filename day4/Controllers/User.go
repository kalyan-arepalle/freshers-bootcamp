//Controllers/User.go
package Controllers

import (
	"day4/Config"
	"day4/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
	"time"
)
var mutex sync.Mutex
var channel chan int
//GetProducts ... Get all products
func GetProducts(c *gin.Context) {
	var prod []Models.Product
	err := Models.GetAllProducts(&prod)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, prod)
	}
}
//CreateProduct ... Create Product
func CreateProduct(c *gin.Context) {
	var prod Models.Product
	err := c.BindJSON(&prod)
	if err != nil {
		return
	}
	err = Models.CreateProduct(&prod)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": prod.Id,
			"product_name": prod.ProductName,
			"price": prod.Price,
			"quantity": prod.Quantity,
			"message": "product successfully added",
		})
	}
}

//GetProductByID ... Get the product by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var prod Models.Product
	err := Models.GetProductByID(&prod, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,prod)
	}
}
//UpdateProduct ... Update the user information
func UpdateProduct(c *gin.Context) {
	var user Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	err = c.BindJSON(&user)
	if err != nil {
		return
	}
	err = Models.UpdateProduct(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteProduct ... Delete the user
func DeleteProduct(c *gin.Context) {
	var prod Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&prod, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func isCoolDownOver(customerId int) bool{
	var order Models.Order
	Config.DB.Model(&order).Where("customer_id = ?",customerId).Last(&order)
	if order.Id == 0{
		return true
	}

	currTime := time.Now()
	diffTime := currTime.Sub(order.CreatedAt).Seconds()

	if diffTime <= 30{
		return false
	}
	return true
}

func isOrderPossible(ord Models.Order, c *gin.Context) bool{
	id := ord.ProductId
	mutex.Lock()
	var prod Models.Product
	err := Models.GetProductByID(&prod,strconv.Itoa(id))
	if err !=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}
	if prod.Quantity < ord.Quantity{
		ord.Status = "Failed"
		mutex.Unlock()
		return false
	}
	prod.Quantity -= ord.Quantity
	ord.Status = "Processed"
	mutex.Unlock()
	err = Models.UpdateProduct(&prod,strconv.Itoa(id))
	if err !=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}
	return true
}

var t *gin.Context
//OrderProduct ... Order product for the user
func OrderProduct(c *gin.Context){
	var order Models.Order
	err := c.BindJSON(&order)
	if err != nil {
		return
	}
	coolDown := isCoolDownOver(order.CustomerId)
	if coolDown == false{
		c.JSON(http.StatusOK,gin.H{
			"message":"Please wait till Cooldown time of 30 seconds",
		})
		return
	}
	possible := isOrderPossible(order,c)
	if possible == false{
		c.JSON(http.StatusOK,gin.H{
			"status":"order failed",
		})
		return
	}
	order.Status = "Order placed"
	err = Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":order.Id,
			"product_id":order.ProductId,
			"quantity":order.Quantity,
			"status":order.Status,
		})
	}
}

//GetOrderByID ... Get the order by id
func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
//GetOrders ... Get all orders
func GetOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//CreateCustomer ... Create Customer
func CreateCustomer(c *gin.Context) {
	var cust Models.Customer
	err := c.BindJSON(&cust)
	if err != nil {
		return
	}
	err = Models.CreateCustomer(&cust)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cust)
	}
}
//GetCustomerByID ... Get the order by id
func GetCustomerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var cust Models.Customer
	err := Models.GetCustomerByID(&cust, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cust)
	}
}
