export interface IDatabaseTable {
    name: string,
    columns?: ITableColumn[],
}

export interface ITableColumn {
    name: string,
    type: string,
}