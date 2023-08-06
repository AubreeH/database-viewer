package mysqlTypes

import (
	"errors"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes/sqlTypes"
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

func GetVariableFromMySqlDataType(t MySqlDataType, isNullable bool) (interface{}, error) {
	if isNullable {
		return getNullableVariableFromMySqlDataType(t)
	}

	return getNonNullableVariableFromMySqlDataType(t)
}

func getNonNullableVariableFromMySqlDataType(t MySqlDataType) (interface{}, error) {
	switch t {
	case CHAR, BINARY, VARCHAR, VARBINARY, TEXT, LONGTEXT, MEDIUMTEXT, TINYTEXT, BLOB, LONGBLOB, MEDIUMBLOB, TINYBLOB, SET, ENUM:
		var s string
		return &s, nil
	case INTEGER, INT, BIGINT, MEDIUMINT, SMALLINT, TINYINT, BIT:
		// var i int64
		var i string
		return &i, nil
	case DATE, DATETIME, TIMESTAMP, TIME, YEAR:
		// var t time.Time
		var t string
		return &t, nil
	case DECIMAL:
		// var f float64
		var f string
		return &f, nil
	case BOOL, BOOLEAN:
		// var b bool
		var b string
		return &b, nil
	}

	return nil, errors.New("unable to find data type " + string(t))
}

func getNullableVariableFromMySqlDataType(t MySqlDataType) (interface{}, error) {
	switch t {
	case CHAR, BINARY, VARCHAR, VARBINARY, TEXT, LONGTEXT, MEDIUMTEXT, TINYTEXT, BLOB, LONGBLOB, MEDIUMBLOB, TINYBLOB, SET, ENUM:
		var s sqlTypes.NullString
		return &s, nil
	case INTEGER, INT, BIGINT, MEDIUMINT, SMALLINT, TINYINT, BIT:
		// var i queryBindingTypes.QBTNullInt64
		var i sqlTypes.NullString
		return &i, nil
	case DATE, DATETIME, TIMESTAMP, TIME, YEAR:
		// var t queryBindingTypes.QBTNullTime
		var t sqlTypes.NullString
		return &t, nil
	case DECIMAL:
		// var f queryBindingTypes.QBTNullFloat64
		var f sqlTypes.NullString
		return &f, nil
	case BOOL, BOOLEAN:
		// var b queryBindingTypes.QBTNullBool
		var b sqlTypes.NullString
		return &b, nil
	}

	return nil, errors.New("unable to find data type " + string(t))
}
