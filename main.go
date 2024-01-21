package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/agunghasbi/schalter-api/middlewares"
	"github.com/agunghasbi/schalter-api/controllers"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")

	router.HandleFunc("/api/v1/events", controllers.GetEvents).Methods("GET")
	router.HandleFunc("/api/v1/events/{id}", controllers.GetEvent).Methods("GET")
	router.HandleFunc("/api/v1/events", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/api/v1/events/{id}", controllers.UpdateEvent).Methods("POST")
	router.HandleFunc("/api/v1/events/{id}", controllers.DeleteEvent).Methods("DELETE")
	
	router.HandleFunc("/api/v1/organizers", controllers.GetOrganizers).Methods("GET")
	router.HandleFunc("/api/v1/organizers/{id}", controllers.GetOrganizer).Methods("GET")
	router.HandleFunc("/api/v1/organizers", controllers.CreateOrganizer).Methods("POST")
	router.HandleFunc("/api/v1/organizers/{id}", controllers.UpdateOrganizer).Methods("POST")
	router.HandleFunc("/api/v1/organizers/{id}", controllers.DeleteOrganizer).Methods("DELETE")
	
	router.HandleFunc("/api/v1/tickets", controllers.GetTickets).Methods("GET")
	router.HandleFunc("/api/v1/tickets/{id}", controllers.GetTicket).Methods("GET")
	router.HandleFunc("/api/v1/events/{eventID}/tickets", controllers.GetTicketsByEventID).Methods("GET")
	router.HandleFunc("/api/v1/tickets", controllers.CreateTicket).Methods("POST")
	router.HandleFunc("/api/v1/tickets/{id}", controllers.UpdateTicket).Methods("POST")
	router.HandleFunc("/api/v1/tickets/{id}", controllers.DeleteTicket).Methods("DELETE")


	// router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	// router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	router.Use(middlewares.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}