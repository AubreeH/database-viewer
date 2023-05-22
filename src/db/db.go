package db

import (
	"changeme/src/connections"

	"github.com/AubreeH/goApiDb/database"
)

var databaseConnections map[string]*database.Database

func GetDatabaseConnection(connection connections.Connection) (*database.Database, bool) {
	db, ok := databaseConnections[connection.Name]
	return db, ok
}

func OpenConnection(connection connections.Connection) error {
	db, err := database.SetupDatabase(database.Config{
		Host:     connection.Host,
		Port:     connection.Port,
		Name:     connection.Database,
		User:     connection.User,
		Password: connection.Password,
	})

	if err != nil {
		return err
	}

	if databaseConnections == nil {
		databaseConnections = make(map[string]*database.Database)
	}

	databaseConnections[connection.Name] = db

	return nil
}

func CloseConnection(connection connections.Connection) error {
	db, ok := GetDatabaseConnection(connection)
	if !ok {
		return nil
	}

	err := db.Db.Close()
	if err != nil {
		return err
	}

	delete(databaseConnections, connection.Name)

	return nil
}
