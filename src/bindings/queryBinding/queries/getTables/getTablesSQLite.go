package getTables

import (
	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes/sqliteTypes"
	"github.com/AubreeH/database-viewer/src/helpers"

	"github.com/AubreeH/goApiDb/access"
	"github.com/AubreeH/goApiDb/database"
)

func getTablesSQLite(dbConnection *database.Database, filter string, itemsPerPage int, offset int) ([]string, error) {
	values, err := access.GetAll(dbConnection, sqliteTypes.SQLiteMaster{}, 0)
	if err != nil {
		return nil, err
	}

	tables := helpers.MapArray(values, func(v sqliteTypes.SQLiteMaster) string {
		return v.Name
	})

	return tables, nil
}
