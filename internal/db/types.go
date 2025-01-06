package db

import (
	"database/sql"
	"time"
)

type CurrencyRate struct {
	CurID           int     `json:"cur_id"`
	Date            string  `json:"date"`
	CurAbbreviation string  `json:"cur_abbreviation"`
	CurScale        int     `json:"cur_scale"`
	CurName         string  `json:"cur_name"`
	CurOfficialRate float64 `json:"cur_officialRate"`
}

type DbInterface interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	//QueryRow(query string, args ...any) *sql.Row
}

type MySQLRepository struct {
	dbConn DbInterface
}

type Repository interface {
	InsertCurrencies(rates []CurrencyRate) error
	GetAllCurrencies() ([]CurrencyRate, error)
	GetCurrenciesDate(date time.Time) ([]CurrencyRate, error)
}
