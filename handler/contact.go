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
	FetchAllContacts(w http.ResponseWriter, r *http.Request)
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

	if payload.Email == nil && payload.PhoneNumber == nil {
		utils.ReturnErrorResponse(w, model.ErrorResponse{
			Message:      "invalid request",
			StatusCode:   http.StatusBadRequest,
			DebugMessage: "email and phone number are null",
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

func (h *contactHandler) FetchAllContacts(w http.ResponseWriter, r *http.Request) {
	var pageInfo model.Page
	err := json.NewDecoder(r.Body).Decode(&pageInfo)
	if err != nil {
		utils.ReturnErrorResponse(w, model.ErrorResponse{
			Message:      "invalid request",
			StatusCode:   http.StatusBadRequest,
			DebugMessage: err.Error(),
		})
		return
	}
	response, err := h.contactService.FetchAllContacts(pageInfo)
	if err != nil {
		utils.ReturnErrorResponse(w, model.ErrorResponse{
			Message:      "error in fetching all contacts",
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
