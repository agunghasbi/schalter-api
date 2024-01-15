package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/agunghasbi/schalter-api/models"
	u "github.com/agunghasbi/schalter-api/utils"
	"net/http"
	"strconv"
	// "time"
)

var CreateEvent = func (w http.ResponseWriter, r *http.Request) {
	event := &models.Event{}

	err := json.NewDecoder(r.Body).Decode(event)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := event.Create()
	u.Respond(w, resp)
}

var UpdateEvent = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id,_ := strconv.ParseUint(params["id"],10,64)
	data := models.GetEvent(id)
	if data == nil {
		u.Respond(w, u.Message(false, "Event not found"))
		return
	}

	event := &models.Event{}
	err := json.NewDecoder(r.Body).Decode(event)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	event.ID = uint(id)
	resp := event.Update()
	u.Respond(w, resp)
}

var DeleteEvent = func (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"],10,64)
	event := models.GetEvent(id)
	if event == nil {
		u.Respond(w, u.Message(false, "Event not found"))
		return
	}
	
	resp := event.Delete()
	u.Respond(w, resp)
}

var GetEvents = func (w http.ResponseWriter, r *http.Request) {

	data := models.GetEvents()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetEvent = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id,_ := strconv.ParseUint(params["id"],10,64)
	data := models.GetEvent(id)
	if data == nil {
		u.Respond(w, u.Message(false, "Event not found"))
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}


