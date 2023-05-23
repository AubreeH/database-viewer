package connectionsbinding

import (
	"changeme/src/connections"
	"changeme/src/db"
	"errors"
)

func (c *ConnectionsBinding) GetTables(connection connections.Connection, filter string) ([]string, error) {
	dbConnection, ok := db.GetDatabaseConnection(connection)
	if !ok {
		return nil, errors.New("not connected")
	}
	rows, err := dbConnection.Db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string

	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}