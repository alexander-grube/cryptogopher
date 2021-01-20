package gopher

import (
	"github.com/alexander-grube/cryptogopher/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Gopher Struct
type Gopher struct {
	gorm.Model
	Name string `json:"name"`
}

// GetGophers returns all Gophers in an array
func GetGophers(c *fiber.Ctx) error {
	db := database.DBConn
	var gophers []Gopher
	db.Find(&gophers)
	return c.JSON(gophers)
}

// GetGopher returns Gopher with given ID
func GetGopher(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var gopher Gopher
	db.Find(&gopher, id)
	return c.JSON(gopher)
}

// NewGopher adds a new Gopher to the DB
func NewGopher(c *fiber.Ctx) error {
	db := database.DBConn
	gopher := new(Gopher)
	if err := c.BodyParser(gopher); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&gopher)
	return c.JSON(gopher)
}

// DeleteGopher deletes the Gopher with given ID from the DB
func DeleteGopher(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var gopher Gopher
	db.First(&gopher, id)
	if gopher.Name == "" {
		return c.Status(500).SendString("No Gopher Found with ID")
	}
	db.Delete(&gopher)
	return c.SendString("Gopher Successfully deleted")
}
