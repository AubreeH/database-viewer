package mysqlTypes

import (
	"database/sql"
	"time"

	"github.com/AubreeH/goApiDb/entities"
)

type ShowColumnsResult struct {
	Field    string
	Type     string
	Nullable string
	Key      sql.NullString
	Default  sql.NullString
	Extra    sql.NullString
}

type InformationSchemaColumns struct {
	entities.EntityBase    `table_name:"information_schema.COLUMNS"`
	TableCatalog           string         `db_name:"TABLE_CATALOG"`
	TableSchema            string         `db_name:"TABLE_SCHEMA"`
	TableName              string         `db_name:"TABLE_NAME"`
	ColumnName             string         `db_name:"COLUMN_NAME"`
	OrdinalPosition        string         `db_name:"ORDINAL_POSITION"`
	ColumnDefault          sql.NullString `db_name:"COLUMN_DEFAULT"`
	IsNullable             string         `db_name:"IS_NULLABLE"`
	DataType               string         `db_name:"DATA_TYPE"`
	CharacterMaximumLength sql.NullInt64  `db_name:"CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   sql.NullInt64  `db_name:"CHARACTER_OCTET_LENGTH"`
	NumericPrecision       sql.NullInt64  `db_name:"NUMERIC_PRECISION"`
	NumericScale           sql.NullInt64  `db_name:"NUMERIC_SCALE"`
	DateTimePrecision      sql.NullInt64  `db_name:"DATETIME_PRECISION"`
	CharacterSetName       sql.NullString `db_name:"CHARACTER_SET_NAME"`
	CollationName          sql.NullString `db_name:"COLLATION_NAME"`
	ColumnType             string         `db_name:"COLUMN_TYPE"`
	ColumnKey              string         `db_name:"COLUMN_KEY"`
	Extra                  string         `db_name:"EXTRA"`
	Privileges             string         `db_name:"PRIVILEGES"`
	ColumnComment          string         `db_name:"COLUMN_COMMENT"`
	IsGenerated            string         `db_name:"IS_GENERATED"`
	GenerationExpression   string         `db_name:"GENERATION_EXPRESSION"`
}

type InformationSchemaKeyColumnUsage struct {
	entities.EntityBase        `table_name:"information_schema.KEY_COLUMN_USAGE"`
	ConstraintCatalog          string         `db_name:"CONSTRAINT_CATALOG"`
	ConstraintSchema           string         `db_name:"CONSTRAINT_SCHEMA"`
	ConstraintName             string         `db_name:"CONSTRAINT_NAME"`
	TableCatalog               string         `db_name:"TABLE_CATALOG"`
	TableSchema                string         `db_name:"TABLE_SCHEMA"`
	TableName                  string         `db_name:"TABLE_NAME"`
	ColumnName                 int64          `db_name:"COLUMN_NAME"`
	OrdinalPosition            int64          `db_name:"ORDINAL_POSITION"`
	PositionInUniqueConstraint sql.NullInt64  `db_name:"POSITION_IN_UNIQUE_CONSTRAINT"`
	ReferencedTableSchema      sql.NullString `db_name:"REFERENCED_TABLE_SCHEMA"`
	ReferencedTableName        sql.NullString `db_name:"REFERENCED_TABLE_NAME"`
	ReferencedColumnName       sql.NullString `db_name:"REFERENCED_COLUMN_NAME"`
}

type InformationSchemaTables struct {
	entities.EntityBase `table_name:"information_schema.TABLES"`
	TableCatalog        string    `db_name:"TABLE_CATALOG"`
	TableSchema         string    `db_name:"TABLE_SCHEMA"`
	TableName           string    `db_name:"TABLE_NAME"`
	TableType           string    `db_name:"TABLE_TYPE"`
	Engine              string    `db_name:"ENGINE"`
	Version             int64     `db_name:"VERSION"`
	RowFormat           string    `db_name:"ROW_FORMAT"`
	TableRows           time.Time `db_name:"TABLE_ROWS"`
	AvgRowLength        int64     `db_name:"AVG_ROW_LENGTH"`
	DataLength          int64     `db_name:"DATA_LENGTH"`
	MaxDataLength       int64     `db_name:"MAX_DATA_LENGTH"`
	IndexLength         int64     `db_name:"INDEX_LENGTH"`
	DataFree            int64     `db_name:"DATA_FREE"`
	AutoIncrement       int64     `db_name:"AUTO_INCREMENT"`
	CreateTime          time.Time `db_name:"CREATE_TIME"`
	UpdateTime          time.Time `db_name:"UPDATE_TIME"`
	CheckTime           time.Time `db_name:"CHECK_TIME"`
	TableCollation      string    `db_name:"TABLE_COLLATION"`
	Checksum            int64     `db_name:"CHECKSUM"`
	CreateOptions       string    `db_name:"CREATE_OPTIONS"`
	TableComment        string    `db_name:"TABLE_COMMENT"`
	MaxIndexLength      int64     `db_name:"MAX_INDEX_LENGTH"`
	Temporary           string    `db_name:"TEMPORARY"`
}
