package getTableColumns

import (
	"errors"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes"
	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/goApiDb/database"
)

func Handle(dbConnection *database.Database, connection connections.Connection, table string) ([]queryBindingTypes.QueryResultColumn, error) {
	switch database.DriverType(connection.Driver) {
	case database.MySql:
		return getTableColumnsMySQL(dbConnection, connection, table)
	case database.SQLite:
		return getTableColumnsSQLite(dbConnection, table)
	}

	return nil, errors.New("driver not supported for GetTableColumns query")
}
