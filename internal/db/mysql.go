package db

import (
	"currency/pkg/mysql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLRepository() (*MySQLRepository, error) {
	conn, err := mysql.MySQLConnectorInit()

	if err != nil {
		return nil, err
	}

	dbRepo := &MySQLRepository{
		dbConn: conn,
	}

	return dbRepo, err
}

// InsertCurrencyRates вставляет список курсов валют в базу данных
func (m *MySQLRepository) InsertCurrencies(rates []CurrencyRate) error {
	query := `
		INSERT INTO currencies (cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate)
		VALUES (?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE cur_official_rate=VALUES(cur_official_rate);
	`

	for _, rate := range rates {
		_, err := m.dbConn.Exec(query, rate.CurID, rate.Date, rate.CurAbbreviation, rate.CurScale, rate.CurName, rate.CurOfficialRate)
		if err != nil {
			return fmt.Errorf("error inserting currency rate: %v", err)
		}
	}

	return nil
}

// GetAllCurrencyRates возвращает все записи курсов валют из базы данных
func (m *MySQLRepository) GetAllCurrencies() ([]CurrencyRate, error) {
	query := `
		SELECT cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate
		FROM currencies;
	`

	rows, err := m.dbConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying currency rates: %v", err)
	}
	defer rows.Close()

	var rates []CurrencyRate
	for rows.Next() {
		var rate CurrencyRate
		err := rows.Scan(&rate.CurID, &rate.Date, &rate.CurAbbreviation, &rate.CurScale, &rate.CurName, &rate.CurOfficialRate)
		if err != nil {
			return nil, fmt.Errorf("error scanning currency rate: %v", err)
		}
		rates = append(rates, rate)
	}

	return rates, nil
}

// GetCurrencyRateByDate возвращает курсы валют за выбранный день
func (m *MySQLRepository) GetCurrenciesDate(date time.Time) ([]CurrencyRate, error) {
	query := `
		SELECT cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate
		FROM currencies
		WHERE date = ?;
	`

	rows, err := m.dbConn.Query(query, date)
	if err != nil {
		return nil, fmt.Errorf("error querying currency rates for date %v: %v", date, err)
	}
	defer rows.Close()

	var rates []CurrencyRate
	for rows.Next() {
		var rate CurrencyRate
		err := rows.Scan(&rate.CurID, &rate.Date, &rate.CurAbbreviation, &rate.CurScale, &rate.CurName, &rate.CurOfficialRate)
		if err != nil {
			return nil, fmt.Errorf("error scanning currency rate: %v", err)
		}
		rates = append(rates, rate)
	}

	return rates, nil
}
