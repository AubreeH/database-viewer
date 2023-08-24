export namespace connections {
	
	export class Connection {
	    name: string;
	    driver: string;
	    driver_name: string;
	    host: string;
	    port: string;
	    database: string;
	    user: string;
	    password?: string;
	    password_save_type: string;
	    sort_order: number;
	    connected: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Connection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.driver = source["driver"];
	        this.driver_name = source["driver_name"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.database = source["database"];
	        this.user = source["user"];
	        this.password = source["password"];
	        this.password_save_type = source["password_save_type"];
	        this.sort_order = source["sort_order"];
	        this.connected = source["connected"];
	    }
	}
	export class PasswordBehaviourDescription {
	    behaviour: string;
	    description: string;
	    display: string;
	
	    static createFrom(source: any = {}) {
	        return new PasswordBehaviourDescription(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.behaviour = source["behaviour"];
	        this.description = source["description"];
	        this.display = source["display"];
	    }
	}

}

export namespace getTableIndexes {
	
	export class IndexResult {
	    database: string;
	    table: string;
	    column: string;
	    // Go type: sql
	    ref_database: any;
	    // Go type: sql
	    ref_table: any;
	    // Go type: sql
	    ref_column: any;
	    position: number;
	    // Go type: sql
	    unique_constraint_position: any;
	
	    static createFrom(source: any = {}) {
	        return new IndexResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.database = source["database"];
	        this.table = source["table"];
	        this.column = source["column"];
	        this.ref_database = this.convertValues(source["ref_database"], null);
	        this.ref_table = this.convertValues(source["ref_table"], null);
	        this.ref_column = this.convertValues(source["ref_column"], null);
	        this.position = source["position"];
	        this.unique_constraint_position = this.convertValues(source["unique_constraint_position"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace queryBindingTypes {
	
	export class QueryResultColumn {
	    field: string;
	    type: string;
	    data_type: string;
	    is_nullable: boolean;
	
	    static createFrom(source: any = {}) {
	        return new QueryResultColumn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.field = source["field"];
	        this.type = source["type"];
	        this.data_type = source["data_type"];
	        this.is_nullable = source["is_nullable"];
	    }
	}
	export class DatabaseColumn {
	    column: QueryResultColumn;
	    indexes: getTableIndexes.IndexResult[];
	    ref_indexes: getTableIndexes.IndexResult[];
	
	    static createFrom(source: any = {}) {
	        return new DatabaseColumn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.column = this.convertValues(source["column"], QueryResultColumn);
	        this.indexes = this.convertValues(source["indexes"], getTableIndexes.IndexResult);
	        this.ref_indexes = this.convertValues(source["ref_indexes"], getTableIndexes.IndexResult);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GetPaginationDetailsResult {
	    total_results: number;
	    total_pages: number;
	    limit: number;
	    offset: number;
	
	    static createFrom(source: any = {}) {
	        return new GetPaginationDetailsResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total_results = source["total_results"];
	        this.total_pages = source["total_pages"];
	        this.limit = source["limit"];
	        this.offset = source["offset"];
	    }
	}
	export class GetTablesResult {
	    tables: string[];
	    details?: GetPaginationDetailsResult;
	
	    static createFrom(source: any = {}) {
	        return new GetTablesResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tables = source["tables"];
	        this.details = this.convertValues(source["details"], GetPaginationDetailsResult);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class QueryResultTablePaginationData {
	    total_results: number;
	    total_pages: number;
	    limit: number;
	    offset: number;
	
	    static createFrom(source: any = {}) {
	        return new QueryResultTablePaginationData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total_results = source["total_results"];
	        this.total_pages = source["total_pages"];
	        this.limit = source["limit"];
	        this.offset = source["offset"];
	    }
	}
	export class QueryResultTableData {
	    rows: any;
	    details: QueryResultTablePaginationData;
	
	    static createFrom(source: any = {}) {
	        return new QueryResultTableData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.rows = source["rows"];
	        this.details = this.convertValues(source["details"], QueryResultTablePaginationData);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

