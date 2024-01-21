package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/agunghasbi/schalter-api/models"
	u "github.com/agunghasbi/schalter-api/utils"
	"net/http"
	"strconv"
)

var CreateTicket = func(w http.ResponseWriter, r *http.Request) {
	ticket := &models.Ticket{}

	err := json.NewDecoder(r.Body).Decode(ticket)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := ticket.Create()
	u.Respond(w, resp)
}

var UpdateTicket = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id,_ := strconv.ParseUint(params["id"],10,64)
	data := models.GetTicket(id)
	if data == nil {
		u.Respond(w, u.Message(false, "Ticket not found"))
		return
	}

	ticket := &models.Ticket{}
	err := json.NewDecoder(r.Body).Decode(ticket)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	ticket.ID = uint(id)
	resp := ticket.Update()
	u.Respond(w, resp)
}

var DeleteTicket = func (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"],10,64)
	ticket := models.GetTicket(id)
	if ticket == nil {
		u.Respond(w, u.Message(false, "Ticket not found"))
		return
	}
	
	resp := ticket.Delete()
	u.Respond(w, resp)
}

var GetTicket = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id,_ := strconv.ParseUint(params["id"],10,64)
	data := models.GetTicket(id)
	if data == nil {
		u.Respond(w, u.Message(false, "Ticket not found"))
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}


var GetTickets = func(w http.ResponseWriter, r *http.Request) {
	// Parse the includeEvent query parameter
	includeEventParam := r.FormValue("includeEvent")
	includeEvent, _ := strconv.ParseBool(includeEventParam)

	data := models.GetTickets(includeEvent)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetTicketsByEventID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["eventID"],10,64)
	data := models.GetTicketsByEventID(uint(id))
	if data == nil {
		u.Respond(w, u.Message(false, "Tickets not found"))
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}