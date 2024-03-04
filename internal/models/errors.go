package models

const (
	CardExpError     = "Your card is expired"
	InvCardNumError  = "Invalid card number"
	InvExpMonthError = "Invalid expiration month"
)

type ServiceError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
