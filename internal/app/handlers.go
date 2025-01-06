package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *App) HandleGetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		currencies, err := app.mysql.GetAllCurrencies()

		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get data from DB: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(currencies)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (app *App) HandleGetCurrenciesByDate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		date, err := parseRequestParams(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get date from params: %v", err), http.StatusBadRequest)
			return
		}

		currencies, err := app.mysql.GetCurrenciesDate(date)

		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get data from DB: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(currencies)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func parseRequestParams(r *http.Request) (time.Time, error) {
	query := r.URL.Query()

	dateParam := query.Get("date")

	date, err := time.Parse("2006-01-02", dateParam)

	if err != nil {
		return time.Now(), err
	}

	location, err := time.LoadLocation("Europe/Minsk")
	if err != nil {
		return time.Now(), err
	}

	return date.In(location).Add(-3 * time.Hour), nil
}
