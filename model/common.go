package model

type SuccessResponse struct {
	StatusCode   int
	DebugMessage string
	Data         interface{}
}

type ErrorResponse struct {
	Message      string `json:"message"`
	StatusCode   int    `json:"-"`
	DebugMessage string `json:"debug_message"`
}
