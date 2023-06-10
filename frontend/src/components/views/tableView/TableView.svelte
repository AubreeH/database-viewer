<script lang="ts">
	import { GetTableData } from "../../../../wailsjs/go/queryBinding/QueryBinding";
	import type { connections, queryBindingTypes } from "../../../../wailsjs/go/models";
	import type { IDatabaseTable } from "../../sidePanels/tableList/types";
	import { setContext } from "svelte";
	import TableHeaderRow from "./subcomponents/TableHeaderRow.svelte";
	import TableBody from "./subcomponents/TableBody.svelte";
	import type { ITableViewContext } from "./types";
	import { writable } from "svelte/store";
	import HiddenScrollbarContainer from "../../common/hiddenScrollbarContainer/HiddenScrollbarContainer.svelte";

	export let connection: connections.Connection;
	export let table: IDatabaseTable;

	let columns: queryBindingTypes.DatabaseColumn[] = undefined;

	const context = writable<ITableViewContext>({
		columns,
		connection,
		table,
	});
	setContext("tableView", context);

	function updateContext(c: ITableViewContext) {
		context.set(c);
	}
	$: updateContext({ columns, connection, table });

	async function getTableData(c: connections.Connection, n: string) {
		if (c && n) {
			columns = await GetTableData(c, n);
		}
	}
	$: getTableData(connection, table.name);
</script>

<HiddenScrollbarContainer>
	<table class="table">
		<TableHeaderRow />
		<TableBody />
	</table>
</HiddenScrollbarContainer>

<style lang="scss">
	.table {
		border-collapse: separate;
		border-spacing: 0;
	}
</style>
