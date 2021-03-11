package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	app.Static("/", "./static/spa", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	app.Static("/gopher", "./services/testdata")
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
	defer database.DBConn.Close()

	setupRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
	//log.Fatal(app.Listen(":8080"))

}
