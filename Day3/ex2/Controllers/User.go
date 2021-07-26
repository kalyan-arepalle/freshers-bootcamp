//Controllers/User.go
package Controllers
import (
	"fmt"
	"net/http"
	"ex2/Models"
	"github.com/gin-gonic/gin"
)

//CreateStudent ... Create Student
func CreateStudent(c *gin.Context) {
	var user Models.Student
	c.BindJSON(&user)
	err := Models.CreateStudent(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//GetStudents ... Get all students
func GetStudents(c *gin.Context) {
	var user []Models.Student
	err := Models.GetAllStudents(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetStudentByID ... Get the Student by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Student
	err := Models.GetStudentByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//UpdateStudent ... Update the Student information
func UpdateStudent(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateStudent(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//DeleteStudent ... Delete the Student
func DeleteStudent(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
