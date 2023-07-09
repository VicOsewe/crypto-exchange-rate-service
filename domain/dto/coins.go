package dto

// Coins represents active coins listed in crypto exchange
type Coins struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}
