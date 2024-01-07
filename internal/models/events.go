package models

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/agunghasbi/schalter-api/database"
)

type Event struct {
	gorm.Model
	OrganizerId int `json:"organizer_id"`
	Name string `json:"organizer_id"`
	DateStart string `json:"organizer_id"`
	DateEnd string `json:"organizer_id"`
	TimeStart string `json:"organizer_id"`
	TimeEnd string `json:"organizer_id"`
	LocationName string `json:"organizer_id"`
	LocationMaps string `json:"organizer_id"`
	Description string `json:"organizer_id"`
	Banner string `json:"organizer_id"`
}

func GetEvents(c *fiber.Ctx){
	db := database.connectDB
	var events []Event
	db.Find(&events)
	c.JSON(events)
}