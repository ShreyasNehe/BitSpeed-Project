package utils

import (
	"encoding/json"
	"github.com/suresh024/identity_reconciliation/model"
	"net/http"
)

// ReturnResponse forms the http response in json format
func ReturnSuccessResponse(w http.ResponseWriter, response model.SuccessResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	en := json.NewEncoder(w)
	_ = en.Encode(response.Data)
}

// ErrorResponse returns generic error response
func ReturnErrorResponse(w http.ResponseWriter, response model.ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	en := json.NewEncoder(w)
	_ = en.Encode(response)
}
