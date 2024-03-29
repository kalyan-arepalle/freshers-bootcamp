//Routes/Routes.go
package Routes
import (
	"ex2/Controllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetStudents)
		grp1.GET("user/:id", Controllers.GetStudentByID)
		grp1.POST("user", Controllers.CreateStudent)
		grp1.PUT("user/:id", Controllers.UpdateStudent)
		grp1.DELETE("user/:id", Controllers.DeleteStudent)
	}
	return r
}