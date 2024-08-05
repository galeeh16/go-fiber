package controller

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/helper"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	// buat variable untuk menampung data user (data ditampung di struct model user)
	var users []entity.User
	// ambil semua data user
	result := database.DB.Find(&users)

	if result.Error != nil {
		panic(result.Error)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    users,
	})
}

// create a user
func CreateUser(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := helper.ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	// insert ke db
	err := database.DB.Create(&newUser).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal membuat user",
			"data":    err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil membuat user",
		"data":    newUser,
	})
}
