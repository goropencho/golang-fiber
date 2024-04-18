package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/goropencho/golang-fibre/database"
	"github.com/goropencho/golang-fibre/lead"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func routes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/{id}", lead.GetLead)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Put("/api/v1/lead/{id}", lead.UpdateLead)
	app.Delete("/api/v1/lead/{id}", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection failed")
	}
	fmt.Println("Connected to Database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Migration Successful")
}

func main() {
	app := fiber.New()
	routes(app)
	initDatabase()
	app.Listen(":3000")
}
