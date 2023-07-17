package usecases

import (
	"log"

	"github.com/VicOsewe/crypto-exchange-rate-service/domain/dao"
	"github.com/VicOsewe/crypto-exchange-rate-service/interfaces/database"
	"github.com/VicOsewe/crypto-exchange-rate-service/interfaces/services"
)

type CryptoExchangeUsecase interface {
	PingCryptoServer() (string, error)
}

// CryptoExchange sets up the crypto exchange usecase layer with all the necessary dependencies
type CryptoExchange struct {
	cryptoServer services.CoinGeckoService
	cryptoDB     database.CryptoDatabase
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

func (c *CryptoExchange) FetchExchangeRateForACryptoAgainstFiat() error {

	coins, err := c.cryptoServer.FetchAvailableCryptocurrencies()
	if err != nil {
		return err
	}

	var coinIDsList []string

	for _, coin := range *coins {
		coinIDsList = append(coinIDsList, coin.ID)
	}

	cryptoPrices, err := c.cryptoServer.FetchExchangeRateForACryptoAgainstFiat(coinIDsList)
	if err != nil {
		return err
	}

	var exchangeRate []dao.ExchangeRate

	for _, cryptoPrice := range *cryptoPrices {
		exchange := &dao.ExchangeRate{
			Cryptocurrency: cryptoPrice.Name,
			USD:            cryptoPrice.USD,
			EURO:           cryptoPrice.EURO,
			GBP:            cryptoPrice.GBP,
		}
		exchangeRate = append(exchangeRate, *exchange)
	}

	err = c.cryptoDB.UpdateExchangeRagesInDB(exchangeRate)
	if err != nil {
		return err
	}
	return nil

}
