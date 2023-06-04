package getTables

import (
	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes/mysqlTypes"
	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/database-viewer/src/helpers"
	"github.com/AubreeH/goApiDb/query"

	"github.com/AubreeH/goApiDb/database"
)

type mySqlQueryResult struct {
	TableName string
}

func getTablesMySQL(dbConnection *database.Database, connection connections.Connection, filter string) ([]string, error) {
	q := query.NewSelectQuery()
	q.Select("t.TABLE_NAME as TableName")
	q.From(mysqlTypes.InformationSchemaTables{}, "t")
	q.Where("t.TABLE_SCHEMA = :databaseName")
	q.SetParam("databaseName", connection.Database)
	result, err := query.ExecuteQuery(dbConnection, q, mySqlQueryResult{})

	if err != nil {
		return nil, err
	}

	tables := helpers.MapArray(result, func(v mySqlQueryResult) string {
		return v.TableName
	})

	return tables, nil
}
