<script lang="ts">
	import { GetTableData, LoadTableData } from "../../../../wailsjs/go/queryBinding/QueryBinding";
	import type { connections, queryBindingTypes } from "../../../../wailsjs/go/models";
	import type { IDatabaseTable } from "../../sidePanels/tableList/types";
	import { setContext } from "svelte";
	import TableHeaderRow from "./subcomponents/TableHeaderRow.svelte";
	import TableBody from "./subcomponents/TableBody.svelte";
	import type { ITableViewContext } from "./types";
	import { writable } from "svelte/store";
	import HiddenScrollbarContainer from "../../common/hiddenScrollbarContainer/HiddenScrollbarContainer.svelte";
	import Icon from "../../common/icon/Icon.svelte";
	import Loader from "../../common/loader/Loader.svelte";

	export let connection: connections.Connection;
	export let table: IDatabaseTable;

	let columns: queryBindingTypes.DatabaseColumn[] = undefined;
	let rows: { [key: string]: any }[] = undefined;
	let loading: boolean = true;

	const context = writable<ITableViewContext>({
		columns,
		connection,
		table,
		rows,
	});
	setContext("tableView", context);

	function updateContext(c: ITableViewContext) {
		context.set(c);
	}
	$: updateContext({ columns, connection, table, rows });

	async function getTableData(c: connections.Connection, n: string) {
		loading = true;
		if (c && n) {
			columns = await GetTableData(c, n);
		}
		loading = false;
	}
	$: getTableData(connection, table.name);

	async function loadTableData(c: connections.Connection, n: string) {
		if (c && n) {
			rows = (await LoadTableData(c, n))?.rows;
		}
	}
	$: loadTableData(connection, table.name);
</script>

{#if loading || true}
	<Loader />
{:else}
	<table class="table">
		<TableHeaderRow />
		<TableBody />
	</table>
{/if}

<style lang="scss">
	.table {
		border-collapse: separate;
		border-spacing: 0;
	}
</style>
