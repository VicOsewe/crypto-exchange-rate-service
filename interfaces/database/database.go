package database

import "github.com/VicOsewe/crypto-exchange-rate-service/domain/dao"

type CryptoDatabase interface {
	UpdateExchangeRagesInDB(rates []dao.ExchangeRate) error
}
