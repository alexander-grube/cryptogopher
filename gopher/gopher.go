package gopher

import (
	"github.com/alexander-grube/cryptogopher/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Gopher struct {
	gorm.Model
	Name string `json:"name"`
}

func GetGophers(c *fiber.Ctx) {
	db := database.DBConn
	var gophers []Gopher
	db.Find(&gophers)
	c.JSON(gophers)
}

func GetGopher(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var gopher Gopher
	db.Find(&gopher, id)
	c.JSON(gopher)
}

func NewGopher(c *fiber.Ctx) {
	db := database.DBConn
	gopher := new(Gopher)
	if err := c.BodyParser(gopher); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&gopher)
	c.JSON(gopher)
}

func DeleteGopher(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var gopher Gopher
	db.First(&gopher, id)
	if gopher.Name == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&gopher)
	c.Send("Gopher Successfully deleted")
}
