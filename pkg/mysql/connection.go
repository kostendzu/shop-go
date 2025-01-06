package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	conn *sql.DB
}

func MySQLConnectorInit() (*MySQL, error) {
	login := os.Getenv("MYSQL_ROOT_USER")
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	connector, err := newMySQLConnector(login, password, host, port, dbName)

	return connector, err
}

func newMySQLConnector(login string, password string, host string, port string, db string) (*MySQL, error) {
	connector := &MySQL{}
	dbConn, err := connector.setConn(login, password, host, port, db)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func (m *MySQL) setConn(login string, password string, host string, port string, db string) (*MySQL, error) {
	if m.conn != nil {
		return m, nil
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		login, password, host, port, db)

	dbConn, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	m.conn = dbConn

	if err := m.conn.Ping(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *MySQL) Query(query string, args ...any) (*sql.Rows, error) {
	return m.conn.Query(query, args...)
}

func (m *MySQL) Exec(query string, args ...any) (sql.Result, error) {
	return m.conn.Exec(query, args...)
}
