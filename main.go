package main

import (
	"github.com/gofiber/fiber"
)

func helloword(c *fiber.Ctx) {
	c.Send("hello,world!")
}
func main() {
	app := fiber.New()
	app.Get("/", helloworld)
	app.Listen(3000)
}
