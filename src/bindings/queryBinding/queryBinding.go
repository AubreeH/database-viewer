package queryBinding

import (
	"context"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queries/getTables"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queries/getTableIndexes"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queries/getTableColumns"
	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes"
	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/database-viewer/src/db"
)

type QueryBinding struct {
	ctx context.Context
}

func NewQueryBinding() (*QueryBinding, func(context.Context)) {
	query := &QueryBinding{}
	return query, func(c context.Context) { query.ctx = c }
}

func (q *QueryBinding) GetTables(connection connections.Connection, filter string) ([]string, error) {
	dbConnection, err := db.GetDatabaseConnection(q.ctx, connection)
	if err != nil {
		return nil, err
	}

	return getTables.Handle(dbConnection, connection, filter)
}

func (q *QueryBinding) GetTableData(connection connections.Connection, table string) ([]queryBindingTypes.DatabaseColumn, error) {
	dbConnection, err := db.GetDatabaseConnection(q.ctx, connection)
	if err != nil {
		return nil, err
	}

	columns, err := getTableColumns.Handle(dbConnection, connection, table)
	if err != nil {
		return nil, err
	}

	indexes, err := getTableIndexes.Handle(dbConnection, connection, table)
	if err != nil {
		return nil, err
	}

	var databaseColumns []queryBindingTypes.DatabaseColumn
	for _, c := range columns {
		column := queryBindingTypes.DatabaseColumn{
			Column: c,
		}
		for _, i := range indexes {
			if i.Column == c.Field {
				if column.Indexes == nil {
					column.Indexes = []getTableIndexes.IndexResult{}
				}
				column.Indexes = append(column.Indexes, i)
			} else if i.RefColumn.Valid && i.RefColumn.String == c.Field {
				if column.RefIndexes == nil {
					column.RefIndexes = []getTableIndexes.IndexResult{}
				}
				column.RefIndexes = append(column.RefIndexes, i)
			}
		}

		databaseColumns = append(databaseColumns, column)
	}

	return databaseColumns, nil
}
