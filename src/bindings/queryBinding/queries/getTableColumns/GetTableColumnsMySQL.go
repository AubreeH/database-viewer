package getTableColumns

import (
	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes"
	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes/mysqlTypes"
	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/goApiDb/database"
	"github.com/AubreeH/goApiDb/query"
)

func getTableColumnsMySQL(dbConnection *database.Database, connection connections.Connection, table string) ([]queryBindingTypes.QueryResultColumn, error) {
	qu := query.NewSelectQuery()
	qu.Select(
		"c.COLUMN_NAME as Field",
		"TRIM(CONCAT(c.COLUMN_TYPE, ' ', c.EXTRA)) as Type",
	)
	qu.From(mysqlTypes.InformationSchemaColumns{}, "c")
	qu.LeftJoin(mysqlTypes.InformationSchemaKeyColumnUsage{}, "kcu", "c.COLUMN_NAME = kcu.COLUMN_NAME AND c.TABLE_NAME = kcu.TABLE_NAME AND c.TABLE_SCHEMA = kcu.TABLE_SCHEMA")
	qu.Where("c.TABLE_NAME = :tableName AND c.TABLE_SCHEMA = :tableSchema")
	qu.SetParam("tableName", table)
	qu.SetParam("tableSchema", connection.Database)
	qu.GroupBy("c.COLUMN_NAME")
	qu.OrderByAscending("c.ORDINAL_POSITION")

	return query.ExecuteQuery(dbConnection, qu, queryBindingTypes.QueryResultColumn{})
}
