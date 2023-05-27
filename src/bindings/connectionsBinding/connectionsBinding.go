package connectionsBinding

import (
	"context"

	"github.com/AubreeH/goApiDb/database"

	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/database-viewer/src/db"
)

type ConnectionsBinding struct {
	ctx context.Context
}

func NewConnectionsBinding() (*ConnectionsBinding, func(context.Context)) {
	connection := &ConnectionsBinding{}
	return connection, func(c context.Context) { connection.ctx = c }
}

func (c *ConnectionsBinding) LoadConnections() ([]connections.Connection, error) {
	return db.GetConnections()
}

func (c *ConnectionsBinding) SaveConnection(connection connections.Connection) error {
	err := connection.Update()
	if err != nil {
		return err
	}

	return db.EmitConnectionsUpdate(c.ctx)
}

func (c *ConnectionsBinding) DeleteConnection(connection connections.Connection) error {
	err := connection.Load(connection.Name)
	if err != nil {
		return err
	}

	err = c.CloseConnection(connection)
	if err != nil {
		return err
	}

	err = connection.Delete()
	if err != nil {
		return err
	}

	return db.EmitConnectionsUpdate(c.ctx)
}

func (c *ConnectionsBinding) OpenConnection(connection connections.Connection) error {
	return db.OpenConnection(c.ctx, connection)
}

func (c *ConnectionsBinding) CloseConnection(connection connections.Connection) error {
	return db.CloseConnection(c.ctx, connection)
}

func (c *ConnectionsBinding) GetPasswordBehaviours() []connections.PasswordBehaviourDescription {
	return connections.GetPasswordBehaviours()
}

func (c *ConnectionsBinding) GetSupportedDrivers() map[string]database.DriverType {
	return map[string]database.DriverType{
		"MySql":      database.MySql,
		"MariaDB":    database.MariaDB,
		"PostgreSQL": database.Postgres,
		"SQLite":     database.SQLite,
	}
}
