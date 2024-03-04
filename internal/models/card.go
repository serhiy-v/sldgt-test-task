package models

type CreditCard struct {
	CardNum  int `json:"Card number"`
	ExpMonth int `json:"Expiration Month"`
	ExpYear  int `json:"Expiration Year"`
}
