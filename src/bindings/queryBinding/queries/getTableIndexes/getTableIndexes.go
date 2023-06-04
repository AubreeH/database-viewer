package getTableIndexes

import (
	"database/sql"
	"errors"

	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/goApiDb/database"
)

type IndexResult struct {
	Database                 string         `json:"database"`
	Table                    string         `json:"table"`
	Column                   string         `json:"column"`
	RefDatabase              sql.NullString `json:"ref_database"`
	RefTable                 sql.NullString `json:"ref_table"`
	RefColumn                sql.NullString `json:"ref_column"`
	Position                 int64          `json:"position"`
	UniqueConstraintPosition sql.NullInt64  `json:"unique_constraint_position"`
}

func Handle(dbConnection *database.Database, connection connections.Connection, table string) ([]IndexResult, error) {
	switch database.DriverType(connection.Driver) {
	case database.MySql:
		return getTableIndexesMySQL(dbConnection, connection, table)
	case database.SQLite:
		return getTableIndexesSQLite(dbConnection, table)
	}

	return nil, errors.New("driver not supported for GetTableIndexes query")
}
