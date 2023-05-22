export namespace connections {
	
	export class Connection {
	    name: string;
	    host: string;
	    port: string;
	    database: string;
	    user: string;
	    password_save_type: string;
	    sort_order: number;
	    connected: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Connection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.database = source["database"];
	        this.user = source["user"];
	        this.password_save_type = source["password_save_type"];
	        this.sort_order = source["sort_order"];
	        this.connected = source["connected"];
	    }
	}

}

