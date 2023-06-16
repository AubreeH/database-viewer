package mysqlTypes

import (
	"database/sql"
	"time"

	"github.com/AubreeH/goApiDb/entities"
)

type ShowColumnsResult struct {
	Field    string         `json:"field"`
	Type     string         `json:"type"`
	Nullable string         `json:"nullable"`
	Key      sql.NullString `json:"key"`
	Default  sql.NullString `json:"default"`
	Extra    sql.NullString `json:"extra"`
}

type InformationSchemaColumns struct {
	entities.EntityBase    `table_name:"information_schema.COLUMNS" json:"-"`
	TableCatalog           string         `db_name:"TABLE_CATALOG" json:"table_catalog"`
	TableSchema            string         `db_name:"TABLE_SCHEMA" json:"table_schema"`
	TableName              string         `db_name:"TABLE_NAME" json:"table_name"`
	ColumnName             string         `db_name:"COLUMN_NAME" json:"column_name"`
	OrdinalPosition        string         `db_name:"ORDINAL_POSITION" json:"ordinal_position"`
	ColumnDefault          sql.NullString `db_name:"COLUMN_DEFAULT" json:"column_default"`
	IsNullable             string         `db_name:"IS_NULLABLE" json:"is_nullable"`
	DataType               string         `db_name:"DATA_TYPE" json:"data_type"`
	CharacterMaximumLength sql.NullInt64  `db_name:"CHARACTER_MAXIMUM_LENGTH" json:"character_maximum_length"`
	CharacterOctetLength   sql.NullInt64  `db_name:"CHARACTER_OCTET_LENGTH" json:"character_octet_length"`
	NumericPrecision       sql.NullInt64  `db_name:"NUMERIC_PRECISION" json:"numeric_precision"`
	NumericScale           sql.NullInt64  `db_name:"NUMERIC_SCALE" json:"numeric_scale"`
	DateTimePrecision      sql.NullInt64  `db_name:"DATETIME_PRECISION" json:"datetime_precision"`
	CharacterSetName       sql.NullString `db_name:"CHARACTER_SET_NAME" json:"character_set_name"`
	CollationName          sql.NullString `db_name:"COLLATION_NAME" json:"collation_name"`
	ColumnType             string         `db_name:"COLUMN_TYPE" json:"column_type"`
	ColumnKey              string         `db_name:"COLUMN_KEY" json:"column_key"`
	Extra                  string         `db_name:"EXTRA" json:"extra"`
	Privileges             string         `db_name:"PRIVILEGES" json:"privileges"`
	ColumnComment          string         `db_name:"COLUMN_COMMENT" json:"column_comment"`
	IsGenerated            string         `db_name:"IS_GENERATED" json:"is_generated"`
	GenerationExpression   string         `db_name:"GENERATION_EXPRESSION" json:"generation_expression"`
}

type InformationSchemaKeyColumnUsage struct {
	entities.EntityBase        `table_name:"information_schema.KEY_COLUMN_USAGE" json:"-"`
	ConstraintCatalog          string         `db_name:"CONSTRAINT_CATALOG" json:"constraint_catalog"`
	ConstraintSchema           string         `db_name:"CONSTRAINT_SCHEMA" json:"constraint_schema"`
	ConstraintName             string         `db_name:"CONSTRAINT_NAME" json:"constraint_name"`
	TableCatalog               string         `db_name:"TABLE_CATALOG" json:"table_catalog"`
	TableSchema                string         `db_name:"TABLE_SCHEMA" json:"table_schema"`
	TableName                  string         `db_name:"TABLE_NAME" json:"table_name"`
	ColumnName                 int64          `db_name:"COLUMN_NAME" json:"column_name"`
	OrdinalPosition            int64          `db_name:"ORDINAL_POSITION" json:"ordinal_position"`
	PositionInUniqueConstraint sql.NullInt64  `db_name:"POSITION_IN_UNIQUE_CONSTRAINT" json:"position_in_unique_constraint"`
	ReferencedTableSchema      sql.NullString `db_name:"REFERENCED_TABLE_SCHEMA" json:"referenced_table_schema"`
	ReferencedTableName        sql.NullString `db_name:"REFERENCED_TABLE_NAME" json:"referenced_table_name"`
	ReferencedColumnName       sql.NullString `db_name:"REFERENCED_COLUMN_NAME" json:"referenced_column_name"`
}

type InformationSchemaTables struct {
	entities.EntityBase `table_name:"information_schema.TABLES" json:"-"`
	TableCatalog        string    `db_name:"TABLE_CATALOG" json:"table_catalog"`
	TableSchema         string    `db_name:"TABLE_SCHEMA" json:"table_schema"`
	TableName           string    `db_name:"TABLE_NAME" json:"table_name"`
	TableType           string    `db_name:"TABLE_TYPE" json:"table_type"`
	Engine              string    `db_name:"ENGINE" json:"engine"`
	Version             int64     `db_name:"VERSION" json:"version"`
	RowFormat           string    `db_name:"ROW_FORMAT" json:"row_format"`
	TableRows           time.Time `db_name:"TABLE_ROWS" json:"table_rows"`
	AvgRowLength        int64     `db_name:"AVG_ROW_LENGTH" json:"avg_row_length"`
	DataLength          int64     `db_name:"DATA_LENGTH" json:"data_length"`
	MaxDataLength       int64     `db_name:"MAX_DATA_LENGTH" json:"max_data_length"`
	IndexLength         int64     `db_name:"INDEX_LENGTH" json:"index_length"`
	DataFree            int64     `db_name:"DATA_FREE" json:"data_free"`
	AutoIncrement       int64     `db_name:"AUTO_INCREMENT" json:"auto_increment"`
	CreateTime          time.Time `db_name:"CREATE_TIME" json:"create_time"`
	UpdateTime          time.Time `db_name:"UPDATE_TIME" json:"update_time"`
	CheckTime           time.Time `db_name:"CHECK_TIME" json:"check_time"`
	TableCollation      string    `db_name:"TABLE_COLLATION" json:"table_collation"`
	Checksum            int64     `db_name:"CHECKSUM" json:"checksum"`
	CreateOptions       string    `db_name:"CREATE_OPTIONS" json:"create_options"`
	TableComment        string    `db_name:"TABLE_COMMENT" json:"table_comment"`
	MaxIndexLength      int64     `db_name:"MAX_INDEX_LENGTH" json:"max_index_length"`
	Temporary           string    `db_name:"TEMPORARY" json:"temporary"`
}
