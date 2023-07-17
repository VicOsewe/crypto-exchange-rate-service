package dao

import "time"

// ExchangeRate represents the exchange rate for between a cryptocurrency and various fiat currencies.
type ExchangeRate struct {
	ID             string    `json:"id"`
	Cryptocurrency string    `json:"cryptocurrency"`
	USD            string    `json:"usd"`
	EURO           string    `json:"euro"`
	GBP            string    `json:"gbp"`
	Timestamp      time.Time `json:"timestamp"`
}
