package queryBinding

import (
	"context"
	"database/sql"

	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/database-viewer/src/db"
	"github.com/AubreeH/goApiDb/database"
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
	rows, err := dbConnection.Db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string

	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}

type DatabaseColumn struct {
	Column QueryResultColumn
	Index  *QueryResultIndex
}

type QueryResultColumn struct {
	Field    string
	Type     string
	Nullable string
	Key      sql.NullString
	Default  sql.NullString
	Extra    sql.NullString
}

type QueryResultIndex struct {
	Table        string
	NonUnique    bool
	KeyName      string
	SeqInIndex   uint
	ColumnName   string
	Collation    string
	Cardinality  uint
	SubPart      sql.NullString
	Packed       sql.NullString
	Null         sql.NullString
	IndexType    string
	Comment      sql.NullString
	IndexComment sql.NullString
}

func (q *QueryBinding) GetTableData(connection connections.Connection, table string) ([]DatabaseColumn, error) {
	dbConnection, err := db.GetDatabaseConnection(q.ctx, connection)
	if err != nil {
		return nil, err
	}

	columns, err := q.getTableColumns(dbConnection, table)
	if err != nil {
		return nil, err
	}

	indexes, err := q.getTableIndexes(dbConnection, table)
	if err != nil {
		return nil, err
	}

	var databaseColumns []DatabaseColumn

	for _, c := range columns {
		column := DatabaseColumn{
			Column: c,
		}
		for _, i := range indexes {
			if i.ColumnName == c.Field {
				column.Index = &i
			}
		}
		databaseColumns = append(databaseColumns, column)
	}

	return databaseColumns, nil
}

func (q *QueryBinding) getTableColumns(dbConnection *database.Database, table string) ([]QueryResultColumn, error) {
	rows, err := dbConnection.Db.Query("SHOW COLUMNS FROM " + table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []QueryResultColumn

	for rows.Next() {
		var column QueryResultColumn
		err := rows.Scan(&column.Field, &column.Type, &column.Nullable, &column.Key, &column.Default, &column.Extra)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	return columns, nil
}

func (c *QueryBinding) getTableIndexes(dbConnection *database.Database, table string) ([]QueryResultIndex, error) {
	rows, err := dbConnection.Db.Query("SHOW INDEXES FROM " + table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var indexes []QueryResultIndex

	for rows.Next() {
		var index QueryResultIndex
		err := rows.Scan(
			&index.Table,
			&index.NonUnique,
			&index.KeyName,
			&index.SeqInIndex,
			&index.ColumnName,
			&index.Collation,
			&index.Cardinality,
			&index.SubPart,
			&index.Packed,
			&index.Null,
			&index.IndexType,
			&index.Comment,
			&index.IndexComment,
		)
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, index)
	}

	return indexes, nil
}

func (c *QueryBinding) getKeyColumnUsage(dbConnection *database.Database, table string, key string) {

}
