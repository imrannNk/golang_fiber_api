package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	UserId      string `json:"UserId"`
	UserName    string `json:"UserName"`
	UserAge     int    `json:"UserAge"`
	UserAddress string `json:"UserAddress"`
}

var users = []User{
	{UserId: "1", UserName: "XXXX", UserAge: 20, UserAddress: "New York"},
	{UserId: "2", UserName: "XXXX", UserAge: 30, UserAddress: "New York"},
	{UserId: "3", UserName: "XXXX", UserAge: 40, UserAddress: "New York"},
	{UserId: "4", UserName: "XXXX", UserAge: 50, UserAddress: "New York"},
}

func get_users(c *fiber.Ctx) error {
	return c.JSON(users)
}

func get_user(c *fiber.Ctx) error {
	userId := c.Params("userId")
	for _, user := range users {
		if user.UserId == userId {
			return c.JSON(user)
		}
	}
	return c.SendStatus(404)
}

func create_user(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.UserId = strconv.Itoa(len(users) + 1)
	users = append(users, user)
	return c.JSON(user)
}

func update_user(c *fiber.Ctx) error {
	userId := c.Params("userId")
	var updatedUser User
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

func delete_user(c *fiber.Ctx) error {
	userId := c.Params("userId")
	for i, user := range users {
		if user.UserId == userId {
			users = append(users[:i], users[i+1:]...)
			return c.SendStatus(204)
		}
	}
	return c.SendStatus(404)
}

func main() {
	app := fiber.New()

	app.Get("/users", get_users)
	app.Get("/users/:userId", get_user)
	app.Post("/users", create_user)
	app.Put("/users/:userId", update_user)
	app.Delete("/users/:userId", delete_user)

	app.Listen(":3000")
}

type Item struct {
	ItemId          string
	ItemName        string
	ItemPrice       float64
	ItemDescription string
}
