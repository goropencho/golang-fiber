package lead

import (
	"github.com/gofiber/fiber/v3"
	"github.com/goropencho/golang-fibre/database"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name" xml:"name" form:"name"`
	Email   string `json:"email" xml:"email" form:"email"`
	Company string `json:"company" xml:"company" form:"company"`
}

func GetLeads(c fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}
func GetLead(c fiber.Ctx) {
	id := c.Params("id")
	var lead Lead
	db := database.DBConn
	db.Find(&lead, id)
	c.JSON(lead)
}

func CreateLead(c fiber.Ctx) {
	lead := new(Lead)
	db := database.DBConn
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(lead, id)
	if lead.Name == "" {
		c.Status(404).Send([]byte("No Leads found"))
	}
	db.Delete(&lead)
	c.Send([]byte("Lead Successfully Deleted"))

}
func UpdateLead(c fiber.Ctx) {

}
