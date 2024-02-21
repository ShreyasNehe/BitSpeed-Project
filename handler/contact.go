package handler

import (
	"encoding/json"
	"github.com/suresh024/identity_reconciliation/model"
	"github.com/suresh024/identity_reconciliation/service"
	"github.com/suresh024/identity_reconciliation/utils"
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
	//funcName := "FetchContacts"

	var payload model.ContactFilter

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.ReturnErrorResponse(w, model.ErrorResponse{
			Message:      "invalid request",
			StatusCode:   http.StatusBadRequest,
			DebugMessage: err.Error(),
		})
		return
	}

	response, err := h.contactService.FetchContacts(payload)
	if err != nil {
		utils.ReturnErrorResponse(w, model.ErrorResponse{
			Message:      "error in fetching contacts",
			StatusCode:   http.StatusBadRequest,
			DebugMessage: err.Error(),
		})
		return
	}

	utils.ReturnSuccessResponse(w, model.SuccessResponse{
		StatusCode:   http.StatusOK,
		DebugMessage: "fetched contacts successfully",
		Data:         response,
	})

}
