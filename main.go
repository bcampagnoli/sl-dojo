package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/sl_dojo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&Person{},
	)

	person := Person{
		Name: "Bruno",
		Age:  44,
	}

	db.Create(&person)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World ??!")
	})

	app.Post("people/", func(c *fiber.Ctx) error {
		var p Person
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		db.Create(&p)
		return c.SendString(fmt.Sprintf("Nome: %s", p.Name))
	})

	log.Fatal(app.Listen(":3000"))
}
