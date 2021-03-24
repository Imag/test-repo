package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main(){
	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/assets", "./assets")
	log.Fatal(app.Listen(":3000"))
}