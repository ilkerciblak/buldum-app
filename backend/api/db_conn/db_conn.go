package dbconn

import (
	"database/sql"
	"fmt"
)

func InitializeSQLDBConnection(driver, conn_str string) (*sql.DB, error) {
	conn, err := sql.Open(driver, conn_str)
	if err != nil {
		return nil, fmt.Errorf("Exception Occured while SQL Connection Openning :\n%v", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("Exception Occured With Pinging SQL DB Connection: \n%v", err)
	}

	return conn, nil
}
