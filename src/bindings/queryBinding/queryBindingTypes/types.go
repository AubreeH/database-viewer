package queryBindingTypes

import (
	"database/sql"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queries/getTableIndexes"
)

type DatabaseColumn struct {
	Column     QueryResultColumn
	Indexes    []getTableIndexes.IndexResult
	RefIndexes []getTableIndexes.IndexResult
}

type QueryResultColumn struct {
	Field string
	Type  string
}

type QueryResultIndex struct {
	Table        string
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

type QueryResultKeyColumnUsage struct {
	ConstrainCatalog           string
	ConstraintSchema           string
	ConstraintName             string
	TableCatalog               string
	TableSchema                string
	TableName                  string
	ColumnName                 string
	OrdinalPosition            int64
	PositionInUniqueConstraint sql.NullInt64
	ReferencedTableSchema      sql.NullString
	ReferencedTableName        sql.NullString
	ReferencedColumnName       sql.NullString
}
