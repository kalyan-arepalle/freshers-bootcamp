//Models/User.go
package Models
import (
	"ex2/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
//GetAllStudents Fetch all Students data
func GetAllStudents(user *[]Student) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}
//CreateStudent ... Insert New data
func CreateStudent(user *Student) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
//GetStudentByID ... Fetch only one user by Id
func GetStudentByID(user *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
//UpdateStudent ... Update user
func UpdateStudent(user *Student, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}
//DeleteStudent ... Delete user
func DeleteStudent(user *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}