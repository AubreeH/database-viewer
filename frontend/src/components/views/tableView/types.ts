export interface ITableViewContext {
    connection: connections.Connection
    table: IDatabaseTable
    columns: queryBindingTypes.DatabaseColumn[]
}