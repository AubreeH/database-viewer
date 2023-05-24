package db

import (
	"github.com/AubreeH/database-viewer/src/connections"

	"github.com/AubreeH/goApiDb/database"
)

var databaseConnections map[string]*database.Database

func GetDatabaseConnection(connection connections.Connection) (*database.Database, error) {
	db, ok := databaseConnections[connection.Name]

	if ok {
		return db, nil
	}

	err := OpenConnection(connection)
	if err != nil {
		return nil, err
	}

	return databaseConnections[connection.Name], nil
}

func GetOpenDatabaseConnection(connection connections.Connection) (*database.Database, bool) {
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
	db, ok := GetOpenDatabaseConnection(connection)
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
