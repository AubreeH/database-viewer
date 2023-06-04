package sqliteTypes

import (
	"database/sql"

	"github.com/AubreeH/goApiDb/entities"
)

type SQLiteMaster struct {
	entities.EntityBase `table_name:"sqlite_master"`
	Type                string `db_name:"type"`
	Name                string `db_name:"name"`
	TableName           string `db_name:"tbl_name"`
	Rootpage            string `db_name:"rootpage"`
	Sql                 string `db_name:"sql"`
}

type PragmaTableInfo struct {
	CID       int64          `db_name:"cid"`
	Name      string         `db_name:"name"`
	Type      string         `db_name:"type"`
	NotNull   bool           `db_name:"notnull"`
	DfltValue sql.NullString `db_name:"dflt_value"`
	PK        int64          `db_name:"pk"`
}
