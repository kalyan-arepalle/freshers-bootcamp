//main.go
package main

import (
	"day4/Config"
	"day4/Models"
	"day4/Routes"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
func main() {
	//channel := make(chan int, 100)
	//Controllers.Initiate(channel)
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