package queryBindingTypes

import (
	"database/sql"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queries/getTableIndexes"
)

type DatabaseColumn struct {
	Column     QueryResultColumn             `json:"column"`
	Indexes    []getTableIndexes.IndexResult `json:"indexes"`
	RefIndexes []getTableIndexes.IndexResult `json:"ref_indexes"`
}

type QueryResultColumn struct {
	Field      string `json:"field"`
	Type       string `json:"type"`
	DataType   string `json:"data_type"`
	IsNullable bool   `json:"is_nullable"`
}

type QueryResultIndex struct {
	Table        string         `json:"table"`
	KeyName      string         `json:"key_name"`
	SeqInIndex   uint           `json:"seq_in_index"`
	ColumnName   string         `json:"column_name"`
	Collation    string         `json:"collation"`
	Cardinality  uint           `json:"cardinality"`
	SubPart      sql.NullString `json:"sub_part"`
	Packed       sql.NullString `json:"packed"`
	Null         sql.NullString `json:"null"`
	IndexType    string         `json:"index_type"`
	Comment      sql.NullString `json:"comment"`
	IndexComment sql.NullString `json:"index_comment"`
}

type QueryResultKeyColumnUsage struct {
	ConstraintCatalog          string         `json:"constraint_catalog"`
	ConstraintSchema           string         `json:"constraint_schema"`
	ConstraintName             string         `json:"constraint_name"`
	TableCatalog               string         `json:"table_catalog"`
	TableSchema                string         `json:"table_schema"`
	TableName                  string         `json:"table_name"`
	ColumnName                 string         `json:"column_name"`
	OrdinalPosition            int64          `json:"ordinal_position"`
	PositionInUniqueConstraint sql.NullInt64  `json:"position_in_unique_constraint"`
	ReferencedTableSchema      sql.NullString `json:"referenced_table_schema"`
	ReferencedTableName        sql.NullString `json:"referenced_table_name"`
	ReferencedColumnName       sql.NullString `json:"referenced_column_name"`
}

type QueryResultTableData struct {
	Rows    any                            `json:"rows"` //[]map[string]interface{}
	Details QueryResultTablePaginationData `json:"details"`
}

type QueryResultTablePaginationData struct {
	TotalResults uint `json:"total_results"`
	TotalPages   uint `json:"total_pages"`
	Limit        uint `json:"limit"`
	Offset       uint `json:"offset"`
}

type GetTablesResult struct {
	Tables  []string                    `json:"tables"`
	Details *GetPaginationDetailsResult `json:"details"`
}

type GetPaginationDetailsResult struct {
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Limit        int `json:"limit"`
	Offset       int `json:"offset"`
}
