import type { connections, queryBindingTypes } from "../../../../wailsjs/go/models";
import type { IDatabaseTable } from "../../sidePanels/tableList/types";

export interface ITableViewContext {
	connection: connections.Connection;
	table: IDatabaseTable;
	columns: queryBindingTypes.DatabaseColumn[];
	rows: { [key: string]: any }[];
	columnSizes: { [key: string]: number }
}
