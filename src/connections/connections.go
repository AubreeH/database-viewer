package connections

import (
	"changeme/src/helpers"
	"encoding/json"
	"errors"
	"os"
)

// PasswordBehaviour describes different methods of behaviour for handling passwords when saving a connection.
// Valid values are SaveAsPlainText, RequestOnConnect, EmptyString
type PasswordBehaviour string

const (
	// SaveAsPlainText - Password is saved as plain text in the connections.json file.
	SaveAsPlainText PasswordBehaviour = "save_as_plain_text"

	// RequestOnConnect - Application will request for the password when attempting to connect to the database.
	RequestOnConnect PasswordBehaviour = "request_on_connect"

	// EmptyString - Application will use an empty string to connect to the database.
	EmptyString PasswordBehaviour = "empty_string"
)

// Connection - The base struct for managing connections for this database viewer client.
type Connection struct {
	Name             string            `json:"name"`
	Host             string            `json:"host"`
	Port             string            `json:"port"`
	Database         string            `json:"database"`
	User             string            `json:"user"`
	Password         string            `json:",omitempty"`
	PasswordSaveType PasswordBehaviour `json:"password_save_type"`
	SortOrder        int               `json:"sort_order"`
	Connected        bool              `json:"connected"`
}

// Connections - Alias type for an array of Connection.
type Connections []Connection

// GetConnections returns all saved Connection values.
func GetConnections() (Connections, error) {
	connections := Connections{}
	err := connections.LoadAll()
	if err != nil {
		return nil, err
	}
	return connections, nil
}

func (connection *Connection) Save() error {
	connections := Connections{}

	if err := connections.LoadAll(); err != nil {
		return err
	}

	if err := connections.add(*connection); err != nil {
		return err
	}

	return connections.saveAll()
}

func (connection *Connection) Update() error {
	connections := Connections{}

	if err := connections.LoadAll(); err != nil {
		return err
	}

	return connections.updateConnection(*connection)
}

func (connection *Connection) Load(name string) error {
	connections := Connections{}
	if err := connections.LoadAll(); err != nil {
		return err
	}

	for _, c := range connections {
		if c.Name == name {
			connection.overwrite(c)
			return nil
		}
	}

	return errors.New("unable to load connection")
}

func (connection *Connection) overwrite(c Connection) {
	connection.Name = c.Name
	connection.Host = c.Host
	connection.Port = c.Port
	connection.User = c.User
	connection.Password = c.Password
	connection.PasswordSaveType = c.PasswordSaveType
	connection.SortOrder = c.SortOrder
}

func (connection *Connection) Delete() error {
	connections := Connections{}
	err := connections.LoadAll()
	if err != nil {
		return err
	}

	for i, c := range connections {
		if c.Name == connection.Name {
			beforeSlice := Connections{}
			if len(connections) > 2 || (len(connections) == 2 && i == 1) {
				beforeSlice = append(beforeSlice, (connections)[:(i)]...)
			}

			if len(connections) > 2 || (len(connections) == 2 && i == 0) {
				beforeSlice = append(beforeSlice, (connections)[(i+1):]...)
			}

			connections = beforeSlice
		}
	}

	return connections.saveAll()
}

func (connections *Connections) LoadAll() error {
	if connections == nil {
		*connections = make(Connections, 1)
	}
	if exists, err := helpers.FileExists("./connections.json"); err != nil {
		return err
	} else if !exists {
		return nil
	}

	data, err := os.ReadFile("./connections.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(data, connections)
}

func (connections *Connections) saveAll() error {
	var connectionsToSave Connections

	for _, v := range *connections {
		connectionsToSave = append(connectionsToSave, v.sanitizeForSaving())
	}
	data, err := json.Marshal(connectionsToSave)
	if err != nil {
		return err
	}

	file, err := os.Create("./connections.json")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func (connections *Connections) add(connection Connection) error {
	for _, c := range *connections {
		if c.Name == connection.Name {
			return errors.New("a connection with this name already exists")
		}
	}

	connection.SortOrder = len(*connections)

	*connections = append(*connections, connection)

	return nil
}

func (connections *Connections) updateConnection(connection Connection) error {
	if !connections.overwriteConnectionInConnections(connection) {
		connections.add(connection)
	}

	return connections.saveAll()
}

func (connections *Connections) overwriteConnectionInConnections(connection Connection) bool {
	for i, c := range *connections {
		if c.Name == connection.Name {
			beforeSlice := Connections{}
			if len(*connections) > 2 || (len(*connections) == 2 && i == 1) {
				beforeSlice = append(beforeSlice, (*connections)[:(i)]...)
			}

			beforeSlice = append(beforeSlice, connection)

			if len(*connections) > 2 || (len(*connections) == 2 && i == 0) {
				beforeSlice = append(beforeSlice, (*connections)[(i+1):]...)
			}

			*connections = beforeSlice
			return true
		}
	}

	return false
}

func (connection *Connection) sanitizeForSaving() Connection {
	name := connection.Name
	host := connection.Host
	port := connection.Port
	user := connection.User
	passwordSaveType := connection.PasswordSaveType
	password := ""
	connected := false
	database := connection.Database

	if passwordSaveType == SaveAsPlainText {
		password = connection.Password
	}

	return Connection{
		Name:             name,
		Host:             host,
		Port:             port,
		User:             user,
		PasswordSaveType: passwordSaveType,
		Password:         password,
		Connected:        connected,
		Database:         database,
	}
}
