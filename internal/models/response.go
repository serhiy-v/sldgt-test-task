package models

type ServiceResponse struct {
	Valid bool `json:"valid"`
}

type ServiceErrorResponse struct {
	Valid bool         `json:"valid"`
	Error ServiceError `json:"error"`
}
