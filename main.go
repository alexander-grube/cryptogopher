package main

import (
	"fmt"

	"github.com/alexander-grube/cryptogopher/database"
	"github.com/alexander-grube/cryptogopher/gopher"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", gopher.GetGophers)
	app.Get("/api/v1/book/:id", gopher.GetGopher)
	app.Post("/api/v1/book", gopher.NewGopher)
	app.Delete("/api/v1/book/:id", gopher.DeleteGopher)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "gophers.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(":3000")

	defer database.DBConn.Close()
}
