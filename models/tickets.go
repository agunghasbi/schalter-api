package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Ticket struct {
	gorm.Model
	EventID uint `json:"event_id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Amount string `json:"amount"`
	Description string `json:"description"`
	SaleStart string `json:"sale_start"`
	SaleEnd string `json:"sale_end"`
}

type TicketWithEvent struct {
	gorm.Model
	EventID uint `json:"event_id"`
	Event Event
	Name string `json:"name"`
	Price string `json:"price"`
	Amount string `json:"amount"`
	Description string `json:"description"`
	SaleStart string `json:"sale_start"`
	SaleEnd string `json:"sale_end"`
}

func GetTickets(preloadEvent bool) ([]*TicketWithEvent) {
	tickets := make([]*TicketWithEvent,0)
	db := GetDB().Table("tickets")
	if preloadEvent {
		db = db.Preload("Event")
	}
	
	err := db.Find(&tickets).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// TODO if preloadEvent is false
	// Remove event from struct

	return tickets
}

func GetTicketsByEventID(eventID uint) ([]*Ticket) {
	tickets := make([]*Ticket,0)
	
	err := GetDB().Table("tickets").Where("event_id = ?",eventID).Find(&tickets).Error;
	if err != nil || len(tickets)==0 {
		fmt.Println(err)
		return nil
	}

	return tickets
}