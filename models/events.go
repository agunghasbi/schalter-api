package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/agunghasbi/schalter-api/utils"
)

type Event struct {
	gorm.Model
	OrganizerId int `json:"organizer_id"`
	Name string `json:"name"`
	DateStart string `json:"date_start"`
	DateEnd string `json:"date_end"`
	TimeStart string `json:"time_start"`
	TimeEnd string `json:"time_end"`
	LocationName string `json:"location_name"`
	LocationMaps string `json:"location_maps"`
	Description string `json:"description"`
	Banner string `json:"banner"`
}

func (event *Event) Validate() (map[string]interface{}, bool){
	if event.Name == "" {
		return u.Message(false,"Event name must be provided"), false
	}
	if event.DateStart == "" {
		return u.Message(false,"Event start date must be provided"), false
	}
	if event.TimeStart == "" {
		return u.Message(false,"Event start time must be provided"), false
	}
	if event.LocationName == "" {
		return u.Message(false,"Event location must be provided"), false
	}
	if event.Description == "" {
		return u.Message(false,"Event description must be provided"), false
	}

	// All the required parameters are present
	return u.Message(true, "success"), true
}

func (event *Event) Create() (map[string]interface{}){
	if resp, ok := event.Validate(); !ok {
		return resp
	}

	if err := GetDB().Create(event).Error; err != nil {
		return u.Message(false, err.Error())
	}

	resp := u.Message(true, "success")
	resp["event"] = event
	return resp
}

func (event *Event) Update() (map[string]interface{}){
	if resp, ok := event.Validate(); !ok {
		return resp
	}

	if err := GetDB().Save(event).Error; err != nil {
		return u.Message(false, err.Error())
	}

	resp := u.Message(true, "success")
	resp["event"] = event
	return resp
}

func (event *Event) Delete() (map[string]interface{}) {
	if err := GetDB().Delete(event).Error; err != nil {
		return u.Message(false, err.Error())
	}

	resp := u.Message(true, "Event successfully deleted.")
	return resp
}

func GetEvents() ([]*Event) {
	events := make([]*Event,0)
	err := GetDB().Table("events").Find(&events).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return events
}

func GetEvent(id uint64) *Event {
	event := &Event{}
	GetDB().Table("events").Where("id = ?", id).First(event)
	if event.Name == "" { //Event not found!
		return nil
	}

	return event
}