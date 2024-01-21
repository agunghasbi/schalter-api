package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/agunghasbi/schalter-api/utils"
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

func (ticket *Ticket) Validate() (map[string]interface{}, bool) {
	if ticket.EventID == 0 {
		return u.Message(false, "Event ID must be provided"), false
	}

	// All the required parameters are present
	return u.Message(true, "success"), true
}

func (ticket *Ticket) Create() (map[string]interface{}) {
	if resp, ok := ticket.Validate(); !ok {
		return resp
	}

	if err := GetDB().Create(ticket).Error; err != nil {
		return u.Message(false, err.Error())
	}

	resp := u.Message(true, "success")
	resp["ticket"] = ticket
	return resp
}

func (ticket *Ticket) Update() (map[string]interface{}) {
	if resp, ok := ticket.Validate(); !ok {
		return resp
	}

	if err := GetDB().Save(ticket).Error; err != nil {
		return u.Message(false, err.Error())
	}

	resp := u.Message(true, "success")
	resp["ticket"] = ticket
	return resp
}

func (ticket *Ticket) Delete() (map[string]interface{}) {
	if err := GetDB().Delete(ticket).Error; err != nil {
		return u.Message(false, err.Error())
	}

	resp := u.Message(true, "Ticket successfully deleted.")
	return resp
}

func GetTicket(id uint64) *Ticket {
	ticket := &Ticket{}
	GetDB().Table("tickets").Where("id = ?", id).First(ticket)
	if ticket.Name == "" { //Ticket not found!
		return nil
	}

	return ticket
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