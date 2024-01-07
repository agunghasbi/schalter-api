package main

import (
	"github.com/agunghasbi/schalter-api/models"
	"github.com/agunghasbi/schalter-api/database"
	"github.com/gofiber/fiber"
)


func setupRoutes(app *fiber.App){
	app.Get("/api/v1/event", book.GetBooks)
	// app.Get("/api/v1/book/:id", book.GetBook)
	// app.Post("/api/v1/book", book.NewBook)
	// app.Delete("/api/v1/book", book.DeleteBook)
}

func main() {
	app := fiber.New()
	// initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}