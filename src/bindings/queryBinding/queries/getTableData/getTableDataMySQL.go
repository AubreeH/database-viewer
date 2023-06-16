package getTableData

import (
	"fmt"

	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queries/getTableColumns"
	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes"
	"github.com/AubreeH/database-viewer/src/bindings/queryBinding/queryBindingTypes/mysqlTypes"
	"github.com/AubreeH/database-viewer/src/connections"
	"github.com/AubreeH/goApiDb/database"
)

func getTableDataMySQL(dbConnection *database.Database, connection connections.Connection, table string) (queryBindingTypes.QueryResultTableData, error) {
	columns, err := getTableColumns.Handle(dbConnection, connection, table)
	if err != nil {
		return queryBindingTypes.QueryResultTableData{}, err
	}

	rows, err := dbConnection.Db.Query(fmt.Sprintf("SELECT * FROM %s LIMIT 100", table))
	if err != nil {
		return queryBindingTypes.QueryResultTableData{}, err
	}

	var output queryBindingTypes.QueryResultTableData
	outputRows := make([]map[string]any, 0)

	columnNames, err := rows.Columns()
	if err != nil {
		return queryBindingTypes.QueryResultTableData{}, err
	}

	for rows.Next() {
		anyRow, interfaceMap, err := constructQueryRow(columns, columnNames)
		if err != nil {
			return queryBindingTypes.QueryResultTableData{}, err
		}

		rows.Scan(anyRow...)

		outputRows = append(outputRows, interfaceMap)
	}

	output.Rows = outputRows

	return output, nil
}

func constructQueryRow(columns []queryBindingTypes.QueryResultColumn, columnsOrder []string) ([]any, map[string]interface{}, error) {
	var anyOutput []any
	interfaceOutput := make(map[string]interface{})

	for _, column := range columns {
		i, err := mysqlTypes.GetVariableFromMySqlDataType(mysqlTypes.MySqlDataType(column.DataType))
		if err != nil {
			return nil, nil, err
		}
		interfaceOutput[column.Field] = i
	}

	for _, column := range columnsOrder {
		anyOutput = append(anyOutput, interfaceOutput[column])
	}

	return anyOutput, interfaceOutput, nil
}
