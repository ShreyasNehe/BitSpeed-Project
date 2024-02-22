package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/suresh024/identity_reconciliation/config"
	"github.com/suresh024/identity_reconciliation/handler"
	"log"
	"net/http"
)

var serviceRoute = "/bitespeed/identity_reconciliation/v1"

func runServer(h handler.Controller) {
	r := mux.NewRouter()

	//health check
	r.HandleFunc(serviceRoute+"/public/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "alive")
	}).Methods("GET")

	r.HandleFunc(serviceRoute+"/contact/identify", h.ContactController.FetchContacts).Methods("POST")
	r.HandleFunc(serviceRoute+"/contact/getAll", h.ContactController.FetchAllContacts).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+config.EnvVariables.EnvPort, r))
}
