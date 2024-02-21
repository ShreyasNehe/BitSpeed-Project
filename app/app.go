package app

import (
	"fmt"
	"github.com/suresh024/identity_reconciliation/config"
	"github.com/suresh024/identity_reconciliation/db"
	"github.com/suresh024/identity_reconciliation/handler"
	"github.com/suresh024/identity_reconciliation/service"
	"github.com/suresh024/identity_reconciliation/store"
	"log"
)

var repos store.Store
var services service.Service
var h handler.Controller

func setupHandlers() {
	h = handler.Controller{
		ContactController: handler.New(services),
	}
}

func setupServices() {
	services = service.Service{
		ContactService: service.NewContactService(repos),
	}
}

func setupRepos() {
	repos = store.Store{
		ContactStore: store.NewContactRepo(),
	}
}

func Start() {
	log.Println("identity_reconciliation - Backend Service starting")
	fmt.Println("identity_reconciliation - Backend Service starting")
	config.InitializeEnv()
	db.InitializeDB()

	setupRepos()
	setupServices()
	setupHandlers()
	runServer(h)
	log.Println("identity_reconciliation - Backend Service started")
	fmt.Println("identity_reconciliation - Backend Service started")
}
