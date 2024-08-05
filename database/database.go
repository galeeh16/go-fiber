package database

import (
	"fmt"
	"go-fiber-gorm/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	// connect to database dengan driver mysql yah
	// https://gorm.io/docs/connecting_to_the_database.html 
	dsn := "root:root@tcp(127.0.0.1:3306)/go-fiber-gorm?charset=utf8mb4&parseTime=True&loc=Local" 
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	helper.PanicIfError(err)

	fmt.Println("Koneksi ke database sukses...")
}