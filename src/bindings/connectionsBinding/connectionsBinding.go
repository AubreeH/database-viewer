package connectionsbinding

import (
	"changeme/src/connections"
	"context"

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
	connections := connections.Connections{}
	if err := connections.LoadAll(); err != nil {
		return nil, err
	}
	return connections, nil
}

func (c *ConnectionsBinding) SaveConnection(connection connections.Connection) error {
	err := connection.Update()
	if err != nil {
		return err
	}

	return c.emitConnectionsUpdate()
}

func (c *ConnectionsBinding) DeleteConnection(name string) error {
	connection := connections.Connection{}
	err := connection.Load(name)
	if err != nil {
		return err
	}
	err = connection.Delete()
	if err != nil {
		return err
	}
	return c.emitConnectionsUpdate()
}

func (c *ConnectionsBinding) emitConnectionsUpdate() error {
	cl, err := connections.GetConnections()
	if err != nil {
		return err
	}
	runtime.EventsEmit(c.ctx, "connections-updated", cl)
	return nil
}
