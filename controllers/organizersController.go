package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	u "github.com/agunghasbi/schalter-api/utils"
	"github.com/agunghasbi/schalter-api/models"
	"net/http"
)

var CreateOrganizer = func (w http.ResponseWriter, r *http.Request) {
	organizer := &models.Organizer{}

	err := json.NewDecoder(r.Body).Decode(organizer)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decodeing request body"))
		return
	}

	resp := organizer.Create()
	u.Respond(w, resp)
	
}
var UpdateOrganizer = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id,_ := strconv.ParseUint(params["id"],10,64)
	data := models.GetOrganizer(uint(id))
	if data == nil {
		u.Respond(w, u.Message(false, "Organizer not found"))
		return
	}

	organizer := &models.Organizer{}
	err := json.NewDecoder(r.Body).Decode(organizer)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	organizer.ID = uint(id)
	resp := organizer.Update()
	u.Respond(w, resp)
}
var DeleteOrganizer = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"],10,64)
	organizer := models.GetOrganizer(uint(id))
	if organizer == nil {
		u.Respond(w, u.Message(false, "Organizer not found"))
		return 
	}

	resp := organizer.Delete()
	u.Respond(w, resp)
}

var GetOrganizers = func (w http.ResponseWriter, r *http.Request) {
	data := models.GetOrganizers()
	resp := u.Message(true,"success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetOrganizer = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"],10,64)
	data := models.GetOrganizer(uint(id))
	if data == nil {
		u.Respond(w, u.Message(false, "Organizer not found"))
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}