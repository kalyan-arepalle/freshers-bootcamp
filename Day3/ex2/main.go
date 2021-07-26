//main.go
package main
import (
	"fmt"
	"github.com/jinzhu/gorm"

	"ex2/Config"
	"ex2/Models"
	"ex2/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})
	//var models = []interface{}{&Models.Student{}, &Models.SubjectMark{}}
	//Config.DB.AutoMigrate(models...)
	r := Routes.SetupRouter()
	//running
	r.Run()
}