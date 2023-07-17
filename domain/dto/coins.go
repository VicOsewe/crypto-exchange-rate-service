package dto

// Coins represents active coins listed in crypto exchange
type Coins struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type CryptoPrices struct {
	Name string `json:"name"`
	USD  string `json:"usd"`
	EURO string `json:"eur"`
	GBP  string `json:"gbp"`
}
