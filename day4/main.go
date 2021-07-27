//main.go
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"day4/Config"
	"day4/Models"
	"day4/Routes"
)

var err error
func main() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())),&gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	//defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{},&Models.Order{},&Models.Customer{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}