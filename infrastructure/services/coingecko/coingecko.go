package coingecko

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/VicOsewe/crypto-exchange-rate-service/configs"
	"github.com/VicOsewe/crypto-exchange-rate-service/domain/dto"
)

const (
	pingUrl      = "/ping"
	coinsListUrl = "/coins/list"
	cryptoPrices = "/simple/price"
)

// RemoteCoinBaseService sets up remote coinbase service with all necessary dependencies
type RemoteCoinBaseService struct {
	URL    string
	client *http.Client
}

// NewRemoteCoinBaseService initializes a new ERP service
func NewRemoteCoinBaseService() *RemoteCoinBaseService {
	r := &RemoteCoinBaseService{
		URL:    configs.MustGetEnvVar("COIN_GECKO_URL"),
		client: &http.Client{},
	}
	r.checkPreconditions()
	return r
}

func (r *RemoteCoinBaseService) checkPreconditions() {
	if r.URL == "" {
		log.Panicf("URL should be defined")
	}
	if r.client == nil {
		log.Panicf("http client not initialized")
	}
}

func (r *RemoteCoinBaseService) makeRequest(
	method string,
	url string,
	body []byte,
) (*http.Response, error) {
	buffer := bytes.NewBuffer(body)
	request, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	return r.client.Do(request)
}

// Ping checks the API status of coingecko
func (r *RemoteCoinBaseService) Ping() (string, error) {
	url := fmt.Sprintf("%s%s", r.URL, pingUrl)
	response, err := r.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("unable to ping coinbase: %w", err)
	}
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read coinbase response: %w", err)
	}
	return string(resp), nil
}

func (r *RemoteCoinBaseService) FetchAvailableCryptocurrencies() (*[]dto.Coins, error) {
	url := fmt.Sprintf("%s%s", r.URL, coinsListUrl)
	response, err := r.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch coin list: %v", err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read coinbase response: %w", err)

	}

	var coinList []dto.Coins
	if err := json.Unmarshal(body, &coinList); err != nil {
		return nil, fmt.Errorf("error unmarshalling coin list: %v", err)
	}
	return &coinList, nil
}

// FetchExchangeRateForACryptoAgainstFiat retrieves the exchange rates for the given cryptocurrency against a list of fiat
// currencies in the url(for now it only retrieves for three USD, GBP,EUR)
func (r *RemoteCoinBaseService) FetchExchangeRateForACryptoAgainstFiat(coinIDsList []string) (*map[string]dto.CryptoPrices, error) {
	cryptoIDs := strings.Join(coinIDsList, ",")
	url := fmt.Sprintf("%s%s?ids=%s&vs_currencies=usd,eur,gbp", r.URL, cryptoPrices, cryptoIDs)
	response, err := r.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}
	var cryptoData map[string]dto.CryptoPrices

	err = json.Unmarshal([]byte(body), &cryptoData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	return &cryptoData, nil
}
