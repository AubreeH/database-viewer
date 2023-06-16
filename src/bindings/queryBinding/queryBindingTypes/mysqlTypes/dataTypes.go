package mysqlTypes

import (
	"errors"
	"time"
)

type MySqlDataType string

const (
	CHAR       MySqlDataType = "char"
	BINARY     MySqlDataType = "binary"
	VARCHAR    MySqlDataType = "varchar"
	VARBINARY  MySqlDataType = "varbinary"
	TEXT       MySqlDataType = "text"
	LONGTEXT   MySqlDataType = "longtext"
	MEDIUMTEXT MySqlDataType = "mediumtext"
	TINYTEXT   MySqlDataType = "tinytext"
	BLOB       MySqlDataType = "blob"
	LONGBLOB   MySqlDataType = "longblob"
	MEDIUMBLOB MySqlDataType = "mediumblob"
	TINYBLOB   MySqlDataType = "tinyblob"
	SET        MySqlDataType = "set"
	ENUM       MySqlDataType = "enum"
	INTEGER    MySqlDataType = "integer"
	INT        MySqlDataType = "int"
	BIGINT     MySqlDataType = "bigint"
	MEDIUMINT  MySqlDataType = "mediumint"
	SMALLINT   MySqlDataType = "smallint"
	TINYINT    MySqlDataType = "tinyint"
	BIT        MySqlDataType = "bit"
	BOOL       MySqlDataType = "bool"
	BOOLEAN    MySqlDataType = "boolean"
	DOUBLE     MySqlDataType = "double"
	DECIMAL    MySqlDataType = "decimal"
	DEC        MySqlDataType = "dec"
	FLOAT      MySqlDataType = "float"
	DATE       MySqlDataType = "date"
	DATETIME   MySqlDataType = "datetime"
	TIMESTAMP  MySqlDataType = "timestamp"
	TIME       MySqlDataType = "time"
	YEAR       MySqlDataType = "year"
)

func GetVariableFromMySqlDataType(t MySqlDataType) (interface{}, error) {
	switch t {
	case CHAR, BINARY, VARCHAR, VARBINARY, TEXT, LONGTEXT, MEDIUMTEXT, TINYTEXT, BLOB, LONGBLOB, MEDIUMBLOB, TINYBLOB, SET, ENUM:
		var s string
		return &s, nil
	case INTEGER, INT, BIGINT, MEDIUMINT, SMALLINT, TINYINT, BIT:
		var i int64
		return &i, nil
	case DATE, DATETIME, TIMESTAMP, TIME, YEAR:
		var t time.Time
		return &t, nil
	case DECIMAL:
		var f float64
		return &f, nil
	case BOOL, BOOLEAN:
		var b bool
		return &b, nil
	}

	return nil, errors.New("unable to find data type " + string(t))
}
