package connectionsBinding

import (
	"database/sql"

	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/database-viewer/src/db"
	"github.com/AubreeH/goApiDb/database"
)

func (c *ConnectionsBinding) GetTables(connection connections.Connection, filter string) ([]string, error) {
	dbConnection, err := db.GetDatabaseConnection(connection)
	if err != nil {
		return nil, err
	}
	c.emitConnectionsUpdate()
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

func (c *ConnectionsBinding) GetTableData(connection connections.Connection, table string) ([]DatabaseColumn, error) {
	dbConnection, err := db.GetDatabaseConnection(connection)
	if err != nil {
		return nil, err
	}
	c.emitConnectionsUpdate()

	columns, err := c.getTableColumns(dbConnection, table)
	if err != nil {
		return nil, err
	}

	indexes, err := c.getTableIndexes(dbConnection, table)
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

func (c *ConnectionsBinding) getTableColumns(dbConnection *database.Database, table string) ([]QueryResultColumn, error) {
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

func (c *ConnectionsBinding) getTableIndexes(dbConnection *database.Database, table string) ([]QueryResultIndex, error) {
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

func (c *ConnectionsBinding) getKeyColumnUsage(dbConnection *database.Database, table string, key string) {

}
