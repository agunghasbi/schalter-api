package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/agunghasbi/schalter-api/utils"
)

type Organizer struct {
	gorm.Model
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}

func (organizer *Organizer) Validate() (map[string]interface{}, bool) {
	if organizer.Name == "" {
		return u.Message(false, "Name must be provided"), false
	}

	return u.Message(true, "success"), true
} 

func (organizer *Organizer) Create() (map[string]interface{}) {
	if resp, ok := organizer.Validate(); !ok {
		return resp
	} 

	if err := GetDB().Create(organizer).Error; err != nil {
		return u.Message(false, err.Error())
	}

	resp := u.Message(true, "success")
	resp["organizer"] = organizer
	return resp
}

func (organizer *Organizer) Update() (map[string]interface{}) {
	if resp, ok := organizer.Validate(); !ok {
		return resp
	}

	if err := GetDB().Save(organizer).Error; err != nil {
		return u.Message(false, err.Error())
	}

	resp := u.Message(true,"success")
	resp["organizer"] = organizer
	return resp
}

func (organizer *Organizer) Delete() (map[string]interface{}) {
	if err := GetDB().Delete(organizer).Error; err != nil {
		return u.Message(false, err.Error())
	}

	resp := u.Message(true, "Organizer successfully deleted")
	return resp
}

func GetOrganizers() ([]*Organizer) {
	organizers := make([]*Organizer,0)
	err := GetDB().Table("organizers").Find(&organizers).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return organizers
}

func GetOrganizer(id uint) *Organizer {
	organizer := &Organizer{}
	GetDB().Table("organizers").Where("id = ?",id).First(organizer)
	if organizer.Name == "" { // Organizer not found
		return nil
	}

	return organizer
}