package db

import (
	"context"

	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/AubreeH/goApiDb/database"
)

var databaseConnections map[string]*database.Database

func GetDatabaseConnection(ctx context.Context, connection connections.Connection) (*database.Database, error) {
	db, ok := databaseConnections[connection.Name]

	if ok {
		return db, nil
	}

	err := OpenConnection(ctx, connection)
	if err != nil {
		return nil, err
	}

	return databaseConnections[connection.Name], nil
}

func GetOpenDatabaseConnection(connection connections.Connection) (*database.Database, bool) {
	db, ok := databaseConnections[connection.Name]
	return db, ok
}

func OpenConnection(ctx context.Context, connection connections.Connection) error {
	db, err := database.SetupDatabase(database.Config{
		Host:     connection.Host,
		Port:     connection.Port,
		Name:     connection.Database,
		User:     connection.User,
		Password: connection.Password,
		Driver:   database.DriverType(connection.Driver),
	})

	if err != nil {
		return err
	}

	if databaseConnections == nil {
		databaseConnections = make(map[string]*database.Database)
	}

	databaseConnections[connection.Name] = db

	return EmitConnectionsUpdate(ctx)
}

func CloseConnection(ctx context.Context, connection connections.Connection) error {
	db, ok := GetOpenDatabaseConnection(connection)
	if !ok {
		return nil
	}

	err := db.Db.Close()
	if err != nil {
		return err
	}

	delete(databaseConnections, connection.Name)

	return EmitConnectionsUpdate(ctx)
}

func EmitConnectionsUpdate(ctx context.Context) error {
	cl, err := GetConnections()
	if err != nil {
		return err
	}

	runtime.EventsEmit(ctx, "connections-updated", cl)

	return nil
}

func GetConnections() (connections.Connections, error) {
	connections, err := connections.GetConnections()
	if err != nil {
		return nil, err
	}

	for i, v := range connections {
		_, ok := GetOpenDatabaseConnection(v)
		connections[i].Connected = ok
	}

	return connections, nil
}
