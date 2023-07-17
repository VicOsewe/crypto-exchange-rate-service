package sql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VicOsewe/crypto-exchange-rate-service/configs"
	"github.com/VicOsewe/crypto-exchange-rate-service/domain/dao"
)

// CryptoExchangeDB sets up cryptoExchange's database layer with ll necessary dependencies
type CryptoExchangeDB struct {
	DB *sql.DB
}

// NewCryptoExchangeDB initializes a new CryptoExchange database instance that meets all the precondition checks
func NewCryptoExchangeDB() *CryptoExchangeDB {
	e := CryptoExchangeDB{
		DB: Init(),
	}
	e.checkPreconditions()
	return &e

}

func (db *CryptoExchangeDB) checkPreconditions() {
	if db.DB == nil {
		log.Panicf("database ORM has not been initialized")
	}
}

func Init() *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.MustGetEnvVar("DB_HOST"),
		configs.MustGetEnvVar("DB_USER"),
		configs.MustGetEnvVar("DB_PASSWORD"),
		configs.MustGetEnvVar("DB_PORT"),
		configs.MustGetEnvVar("DB_NAME"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (db *CryptoExchangeDB) UpdateExchangeRagesInDB(rates []dao.ExchangeRate) error {
	// Prepare the SQL statement
	stmt, err := db.DB.Prepare(`
		INSERT INTO exchange_rates (cryptocurrency, fiat, rate, timestamp)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (cryptocurrency, fiat) DO UPDATE
		SET rate = EXCLUDED.rate, timestamp = EXCLUDED.timestamp
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement in a transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	for _, rate := range rates {
		_, err := tx.Stmt(stmt).Exec(rate.Cryptocurrency, rate.USD, rate.GBP, rate.EURO, rate.Timestamp)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
