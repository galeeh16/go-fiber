package migration

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/helper"
	"go-fiber-gorm/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	helper.PanicIfError(err)

	fmt.Println("Database di migrasi...")
}