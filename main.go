package main

import (
	"fmt"

	config "go-api-sql/Config"
	models "go-api-sql/Models"
	routes "go-api-sql/Routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})
	r := routes.SetupRouter()
	//running
	r.Run("localhost:4040")
}
