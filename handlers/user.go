package handlers

import (
	"strconv"

	"go-fiber-api/models"

	"github.com/gofiber/fiber/v2"
)

var users = []models.User{
	{UserId: "1", UserName: "XXXX", UserAge: 20, UserAddress: "New York"},
	{UserId: "2", UserName: "XXXX", UserAge: 30, UserAddress: "New York"},
	{UserId: "3", UserName: "XXXX", UserAge: 40, UserAddress: "New York"},
	{UserId: "4", UserName: "XXXX", UserAge: 50, UserAddress: "New York"},
}

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	for _, user := range users {
		if user.UserId == userId {
			return c.JSON(user)
		}
	}
	return c.SendStatus(404)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.UserId = strconv.Itoa(len(users) + 1)
	users = append(users, user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	var updatedUser models.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return err
	}
	for i, user := range users {
		if user.UserId == userId {
			users[i] = updatedUser
			return c.JSON(updatedUser)
		}
	}
	return c.SendStatus(404)
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	for i, user := range users {
		if user.UserId == userId {
			users = append(users[:i], users[i+1:]...)
			return c.SendStatus(204)
		}
	}
	return c.SendStatus(404)
}
