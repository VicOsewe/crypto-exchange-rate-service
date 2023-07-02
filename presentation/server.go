package presentation

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/VicOsewe/crypto-exchange-rate-service/infrastructure/services/coingecko"
	rest "github.com/VicOsewe/crypto-exchange-rate-service/presentation/http/handlers"
	"github.com/VicOsewe/crypto-exchange-rate-service/usecases"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	serverTimeoutSeconds = 120
)

// InitHandlers initializes all the handlers dependencies
func InitHandlers() *rest.RestfulAPIs {
	cryptoService := coingecko.NewRemoteCoinBaseService()
	usecases := usecases.NewCryptoExchange(cryptoService)
	return rest.NewRestfulPIs(usecases)
}

func Router() (*mux.Router, error) {
	r := mux.NewRouter()
	h := InitHandlers()
	RESTRoutes := r.PathPrefix("/api/v1/").Subrouter()
	RESTRoutes.Path("/ping").Methods(http.MethodGet, http.MethodOptions).HandlerFunc(h.PingCryptoServer())
	return r, nil
}

func PrepareServer(port int) *http.Server {
	r, err := Router()
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
	addr := fmt.Sprintf(":%d", port)
	h := handlers.CompressHandlerLevel(r, gzip.BestCompression)

	h = handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
	)(h)
	h = handlers.CombinedLoggingHandler(os.Stdout, h)
	h = handlers.ContentTypeHandler(
		h,
		"application/json",
		"application/x-www-form-urlencoded",
	)
	return &http.Server{
		Handler:      h,
		Addr:         addr,
		WriteTimeout: serverTimeoutSeconds * time.Second,
		ReadTimeout:  serverTimeoutSeconds * time.Second,
	}
}
