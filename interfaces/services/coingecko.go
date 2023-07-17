package services

import "github.com/VicOsewe/crypto-exchange-rate-service/domain/dto"

// CoinGeckoService represents's cryptocurrency data acquisition level interfaces
type CoinGeckoService interface {
	Ping() (string, error)
	FetchAvailableCryptocurrencies() (*[]dto.Coins, error)
	FetchExchangeRateForACryptoAgainstFiat(coinIDsList []string) (*map[string]dto.CryptoPrices, error)
}
