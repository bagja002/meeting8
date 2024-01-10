package main

import (
	"belajar-database/controllers"
	"belajar-database/database"
	"log"

	//"belajar-database/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	database.Connect()

    app.Get("/getDataall",controllers.GetDataAll)
	app.Post("/createData", controllers.CreateData)
	app.Get("/getOneData", controllers.GetOneData)
	app.Post("/UpdateData",controllers.UpdateData)

    log.Fatal(app.Listen(":4000"))
}