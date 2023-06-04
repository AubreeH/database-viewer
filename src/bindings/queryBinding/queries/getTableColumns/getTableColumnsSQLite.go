package getTableColumns

import (
	"fmt"

	"github.com/AubreeH/goApiDb/query"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes"
	"github.com/AubreeH/goApiDb/database"
)

func getTableColumnsSQLite(dbConnection *database.Database, table string) ([]queryBindingTypes.QueryResultColumn, error) {
	q := query.NewSelectQuery()
	q.Select("p.name", "p.type")
	q.FromTable(fmt.Sprintf("pragma_table_info('%s')", table), "p")
	result, err := query.ExecuteQuery(dbConnection, q, queryBindingTypes.QueryResultColumn{})
	if err != nil {
		return nil, err
	}

	return result, nil
}
