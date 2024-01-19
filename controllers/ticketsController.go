package controllers

import (
	"github.com/gorilla/mux"
	"github.com/agunghasbi/schalter-api/models"
	u "github.com/agunghasbi/schalter-api/utils"
	"net/http"
	"strconv"
)

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