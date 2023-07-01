package dao

import "time"

// ExchangeRate represents the exchange rate for between a cryptocurrency and a fiat currency.
type ExchangeRate struct {
	ID             string    `json:"id"`
	Cryptocurrency string    `json:"cryptocurrency"`
	Fiat           float64   `json:"fiat"`
	Rate           float64   `json:"rate"`
	Timestamp      time.Time `json:"timestamp"`
}
