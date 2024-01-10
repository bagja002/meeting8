package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"belajar-database/models"
)

var DB *gorm.DB

func Connect(){

	dsn := "root@tcp(127.0.0.1:3306)/meeting7?charset=utf8mb4&parseTime=True&loc=Local"
  	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil {
		panic("Tidak Terhubung")
	}

	fmt.Println("Database Telah Terhubung ")

	DB = conn

	conn.AutoMigrate(models.Admin{}, models.Users{})

}
