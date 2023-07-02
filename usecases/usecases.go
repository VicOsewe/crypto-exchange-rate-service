package usecases

import (
	"log"

	"github.com/VicOsewe/crypto-exchange-rate-service/interfaces/services"
)

type CryptoExchangeUsecase interface {
	PingCryptoServer() (string, error)
}

// CryptoExchange sets up the crypto exchange usecase layer with all the necessary dependencies
type CryptoExchange struct {
	cryptoServer services.CoinGeckoService
}

// NewCryptoExchange initializes a crypto exchange usecase instance that meets all the precondition checks
func NewCryptoExchange(
	crypto services.CoinGeckoService,
) *CryptoExchange {
	c := &CryptoExchange{
		cryptoServer: crypto,
	}
	c.checkPreconditions()
	return c
}

func (c *CryptoExchange) checkPreconditions() {
	if c.cryptoServer == nil {
		log.Panicf("usecase: crypto exchange service is not initialized")
	}
}

// Ping checks if the crypto server of choice is connected
func (c *CryptoExchange) PingCryptoServer() (string, error) {
	return c.cryptoServer.Ping()
}
