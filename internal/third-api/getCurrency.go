package third_api

import (
	"currency/internal/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiURL = "https://api.nbrb.by/exrates/rates?periodicity=0"

func FetchCurrencyRates() ([]db.CurrencyRate, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from NBRB API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var rates []db.CurrencyRate
	err = json.Unmarshal(body, &rates)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return rates, nil
}
