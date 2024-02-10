package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// get
func home(c *fiber.Ctx) error {
	return c.Render("addNoteForm", fiber.Map{"Title": "Home", "Content": "addNoteForm"}, "layout")
}

func signup(c *fiber.Ctx) error {
	return c.Render("signupForm", fiber.Map{"Title": "Sign Up", "Content": "signupForm"}, "layout")
}

// post
func addUser(c *fiber.Ctx) error {
	newUser := new(User)

	if err := c.BodyParser(newUser); err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	log.Println(newUser.Username, "added")

	return c.SendStatus(fiber.StatusOK)
}

func addNote(c *fiber.Ctx) error {
	newNote := new(Todo)

	if err := c.BodyParser(newNote); err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	return c.SendString(fmt.Sprintf("<li>%s</li>", newNote.Content))
}
