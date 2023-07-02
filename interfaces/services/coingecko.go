package services

// CoinGeckoService represents's cryptocurrency data acquisition level interfaces
type CoinGeckoService interface {
	Ping() (string, error)
}
