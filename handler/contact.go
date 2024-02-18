package handler

import (
	"github.com/suresh024/identity_reconciliation/service"
	"net/http"
)

type ContactHandler interface {
	FetchContacts(w http.ResponseWriter, r *http.Request)
}
type contactHandler struct {
	contactService service.ContactService
}

func New(service service.Service) ContactHandler {
	return &contactHandler{
		contactService: service.ContactService,
	}
}

func (h *contactHandler) FetchContacts(w http.ResponseWriter, r *http.Request) {

}
