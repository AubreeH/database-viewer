package connectionsbinding

import (
	"changeme/src/connections"
	"changeme/src/db"
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ConnectionsBinding struct {
	ctx context.Context
}

func NewConnectionsBinding() (*ConnectionsBinding, func(context.Context)) {
	connection := &ConnectionsBinding{}
	return connection, func(c context.Context) { connection.ctx = c }
}

func (c *ConnectionsBinding) LoadConnections() ([]connections.Connection, error) {
	return c.getConnections()
}

func (c *ConnectionsBinding) SaveConnection(connection connections.Connection) error {
	err := connection.Update()
	if err != nil {
		return err
	}

	return c.emitConnectionsUpdate()
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

	return c.emitConnectionsUpdate()
}

func (c *ConnectionsBinding) OpenConnection(connection connections.Connection) error {
	err := db.OpenConnection(connection)
	runtime.LogDebug(c.ctx, fmt.Sprint(connection))
	if err != nil {
		return nil
	}
	fmt.Println("test123")
	return c.emitConnectionsUpdate()
}

func (c *ConnectionsBinding) CloseConnection(connection connections.Connection) error {
	err := db.CloseConnection(connection)
	if err != nil {
		return nil
	}
	return c.emitConnectionsUpdate()
}

func (c *ConnectionsBinding) emitConnectionsUpdate() error {
	cl, err := c.getConnections()
	if err != nil {
		return err
	}

	runtime.EventsEmit(c.ctx, "connections-updated", cl)

	return nil
}

func (c *ConnectionsBinding) getConnections() (connections.Connections, error) {
	connections, err := connections.GetConnections()
	if err != nil {
		return nil, err
	}

	for i, v := range connections {
		_, ok := db.GetDatabaseConnection(v)
		connections[i].Connected = ok
	}

	return connections, nil
}
