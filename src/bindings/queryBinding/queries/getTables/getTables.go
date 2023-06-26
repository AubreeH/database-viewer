package getTables

import (
	"errors"

	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/goApiDb/database"
)

func Handle(dbConnection *database.Database, connection connections.Connection, filter string, offset int) ([]string, error) {
	switch database.DriverType(connection.Driver) {
	case database.MySql:
		return getTablesMySQL(dbConnection, connection, filter, offset)
	case database.SQLite:
		return getTablesSQLite(dbConnection, filter)
	}

	return nil, errors.New("driver not supported for GetTables query")
}
