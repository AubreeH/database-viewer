package getTableData

import (
	"errors"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes"
	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/goApiDb/database"
)

func Handle(dbConnection *database.Database, connection connections.Connection, table string, limit, offset uint) (queryBindingTypes.QueryResultTableData, error) {
	switch database.DriverType(connection.Driver) {
	case database.MySql:
		return getTableDataMySQL(dbConnection, connection, table, limit, offset)
	case database.SQLite:
		return getTableDataSQLite(dbConnection, connection, table, limit, offset)
	}

	return queryBindingTypes.QueryResultTableData{}, errors.New("driver not supported for GetTableData query")
}
