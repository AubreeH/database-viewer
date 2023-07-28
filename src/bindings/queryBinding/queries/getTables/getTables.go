package getTables

import (
	"errors"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes"
	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/goApiDb/database"
)

func Handle(dbConnection *database.Database, connection connections.Connection, filter string, offset int) ([]string, *queryBindingTypes.GetPaginationDetailsResult, error) {
	itemsPerPage := 25
	switch database.DriverType(connection.Driver) {
	case database.MySql:
		return getTablesMySQL(dbConnection, connection, filter, itemsPerPage, offset)
	case database.SQLite:
		result, err := getTablesSQLite(dbConnection, filter, itemsPerPage, offset)
		return result, nil, err
	}

	return nil, nil, errors.New("driver not supported for GetTables query")
}
