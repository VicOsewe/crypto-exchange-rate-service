package usecases

import (
	"fmt"
	"log"

	"github.com/VicOsewe/crypto-exchange-rate-service/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// CryptoExchangeDB sets up cryptoExchange's database layer with ll necessary dependencies
type CryptoExchangeDB struct {
	DB *gorm.DB
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

func runMigrations(db *gorm.DB) {
	tables := []interface{}{}
	for _, table := range tables {
		if err := db.AutoMigrate(table); err != nil {
			log.Panicf("can't run migrations on table %s: err: %v", table, err)
		}
	}
}

// Init initializes a new gorm DB instance by connecting to the database specified.
func Init() *gorm.DB {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.MustGetEnvVar("DB_HOST"),
		configs.MustGetEnvVar("DB_USER"),
		configs.MustGetEnvVar("DB_PASSWORD"),
		configs.MustGetEnvVar("DB_PORT"),
		configs.MustGetEnvVar("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("can't open connection to the local database: %v", err)
	}
	runMigrations(db)
	return db
}
