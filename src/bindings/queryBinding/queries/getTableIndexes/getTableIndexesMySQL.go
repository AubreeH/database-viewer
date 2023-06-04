package getTableIndexes

import (
	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes/mysqlTypes"
	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/goApiDb/database"
	"github.com/AubreeH/goApiDb/query"
)

func getTableIndexesMySQL(dbConnection *database.Database, connection connections.Connection, table string) ([]IndexResult, error) {
	q := query.NewSelectQuery()
	q.Select(
		"kcu.TABLE_SCHEMA as Db",
		"kcu.TABLE_NAME as Ta",
		"kcu.COLUMN_NAME as Co",
		"kcu.REFERENCED_TABLE_SCHEMA as RefDatabase",
		"kcu.REFERENCED_TABLE_NAME as RefTable",
		"kcu.REFERENCED_COLUMN_NAME as RefColumn",
		"kcu.ORDINAL_POSITION as Position",
		"kcu.POSITION_IN_UNIQUE_CONSTRAINT as UniqueConstraintPosition",
	)
	q.From(mysqlTypes.InformationSchemaKeyColumnUsage{}, "kcu")
	q.Where("(kcu.TABLE_NAME = :tableName AND kcu.TABLE_SCHEMA = :tableSchema) OR (kcu.REFERENCED_TABLE_NAME = :tableName AND kcu.REFERENCED_TABLE_SCHEMA = :tableSchema)")
	q.SetParam("tableName", table)
	q.SetParam("tableSchema", connection.Database)

	return query.ExecuteQuery(dbConnection, q, IndexResult{})
}
