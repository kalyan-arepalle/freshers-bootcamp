//Routes/Routes.go
package Routes
import (
	"github.com/gin-gonic/gin"

	"day4/Controllers"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	product := r.Group("/product")
	{
		product.GET("", Controllers.GetProducts)
		product.POST("", Controllers.CreateProduct)
		product.GET(":id", Controllers.GetProductByID)
		product.PATCH(":id", Controllers.UpdateProduct)
		product.DELETE(":id", Controllers.DeleteProduct)
	}

	order :=r.Group("/order")
	{
		order.POST("",Controllers.OrderProduct)
		order.GET(":id",Controllers.GetOrderByID)
		order.GET("",Controllers.GetOrders)
	}

	customer :=r.Group("/customer")
	{
		customer.POST("",Controllers.CreateCustomer)
		customer.GET(":id",Controllers.GetCustomerByID)
	}
	return r
}