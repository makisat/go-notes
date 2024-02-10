package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Todo struct {
	Content string `json:"content"`
}

type User struct {
	Username string `json:"username"`
	Hashed   string `json:"hashed"`
	Todo     []Todo `json:"todo"`
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// get
	app.Get("/home", home)
	app.Get("/signup", signup)

	// post
	app.Post("/addUser", addUser)
	app.Post("/addNote", addNote)

	log.Fatal(app.Listen(":8080"))
}
