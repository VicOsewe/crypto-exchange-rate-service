package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VicOsewe/crypto-exchange-rate-service/domain/dto"
	"github.com/VicOsewe/crypto-exchange-rate-service/usecases"
)

// Presentation represents the presentation layer contract
// type Presentation interface {
// }

// RestfulAPIs sets up RESTFUL APIs with all necessary dependencies
type RestfulAPIs struct {
	usecases usecases.CryptoExchangeUsecase
}

// NewRestfulAPIS initializes a new REST implementation
func NewRestfulPIs(usecases usecases.CryptoExchangeUsecase) *RestfulAPIs {
	rst := &RestfulAPIs{
		usecases: usecases,
	}
	rst.checkPreconditions()
	return rst
}

func (rst *RestfulAPIs) checkPreconditions() {
	if rst.usecases == nil {
		log.Panicf("usecases not initialized")
	}
}

// PingCryptoServer
func (rst *RestfulAPIs) PingCryptoServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := rst.usecases.PingCryptoServer()
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		marshalled, err := json.Marshal(dto.APIResponseMessage{
			Message:    "server pinged successfully",
			StatusCode: http.StatusOK,
			Body:       response,
		})
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		RespondWithJSON(w, http.StatusOK, marshalled)
	}
}
