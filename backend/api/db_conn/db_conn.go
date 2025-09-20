package dbconn

import (
	"database/sql"
	"fmt"
)

type SqlConnectionConfig struct {
	Driver           string
	ConnectionString string
}

func NewSqlConnectionConfig(driver, conn_str string) *SqlConnectionConfig {
	return &SqlConnectionConfig{
		Driver:           driver,
		ConnectionString: conn_str,
	}
}

func (s *SqlConnectionConfig) InitializeSQLDBConnection(errChan chan<- error) *sql.DB {

	conn, err := sql.Open(s.Driver, s.ConnectionString)
	if err != nil {
		errChan <- fmt.Errorf("Error Occured While sql.Open with %v", err)
		return nil
	}

	if err := conn.Ping(); err != nil {
		errChan <- fmt.Errorf("Error Occured While conn.Ping with %v", err)
		return nil
	}

	return conn
}
