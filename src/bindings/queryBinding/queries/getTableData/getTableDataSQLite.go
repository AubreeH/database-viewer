package getTableData

import (
	"errors"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes"
	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/goApiDb/database"
)

func getTableDataSQLite(dbConnection *database.Database, connection connections.Connection, table string) (queryBindingTypes.QueryResultTableData, error) {
	return queryBindingTypes.QueryResultTableData{}, errors.New("getTableData: sqlite not supported")
}
