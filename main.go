package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string
	Role     string
}

func main() {
	app := fiber.New()

	// Routes.
	app.Get("/post", handleGetPost)                         //public
	app.Get("/post/manage", onlyAdmin(handleGetPostManage)) // admin
	app.Get("/post/ex", onlyEx(handleGetPostEx))            // admin

	log.Fatal(app.Listen(":8080"))
}

// Auth on admin role.
func onlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Here we use the user we created.
		user := getUserFromDB()
		if user.Role != "admin" {
			return c.SendStatus(http.StatusUnauthorized)
		}
		return fn(c)
	}
}

// Auth on two roles.
func onlyEx(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := getUserFromDB()
		if user.Role == "admin" || user.Role == "ex" {
			return fn(c)
		} else {
			return c.SendStatus(http.StatusUnauthorized)
		}
	}
}

func getUserFromDB() User {
	return User{
		Username: "Tech",
		Role:     "ex",
	}
}

// Handlers..
func handleGetPost(c *fiber.Ctx) error {
	return c.JSON("Some posts here")
}

func handleGetPostManage(c *fiber.Ctx) error {
	return c.JSON("ADMIN PAGE")
}

func handleGetPostEx(c *fiber.Ctx) error {
	return c.JSON("exxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}
