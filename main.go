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
	app.Get("/api/v1/gopher", gopher.GetGophers)
	app.Get("/api/v1/gopher/:id", gopher.GetGopher)
	app.Post("/api/v1/gopher", gopher.NewGopher)
	app.Delete("/api/v1/gopher/:id", gopher.DeleteGopher)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "gophers.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&gopher.Gopher{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	// port := os.Getenv("PORT")
	// app.Listen(":" + port)
	app.Listen(":8080")

	defer database.DBConn.Close()
}
