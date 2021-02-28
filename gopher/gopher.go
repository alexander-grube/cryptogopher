package gopher

import (
	"crypto/rand"
	"encoding/base64"
	"strings"

	"github.com/alexander-grube/cryptogopher/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Gopher Struct
type Gopher struct {
	gorm.Model
	Name string `json:"name"`
	Seed string `json:"seed"`
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
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
	if gopher.Name == "" {
		return c.Status(404).SendString("No Gopher Found with ID")
	}
	return c.JSON(gopher)
}

// NewGopher adds a new Gopher to the DB
func NewGopher(c *fiber.Ctx) error {
	db := database.DBConn
	gopher := new(Gopher)

	if err := c.BodyParser(gopher); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	if valid, invalidFields := checkGopher(gopher); !valid {
		return c.Status(400).SendString("Invalid Input(" + invalidFields + ")")
	}

	var err error
	gopher.Seed, err = GenerateRandomString(32)
	if err != nil {
		return c.Status(501).SendString(err.Error())
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

func checkGopher(gopher *Gopher) (bool, string) {
	var errors string
	if len(strings.TrimSpace(gopher.Name)) == 0 {
		errors += "Name, "
		return false, errors
	}
	return true, errors
}
