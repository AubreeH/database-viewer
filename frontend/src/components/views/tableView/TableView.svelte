<script lang="ts">
	import { GetTableData, LoadTableData } from "../../../../wailsjs/go/queryBinding/QueryBinding";
	import type { connections, queryBindingTypes } from "../../../../wailsjs/go/models";
	import type { IDatabaseTable } from "../../sidePanels/tableList/types";
	import { setContext } from "svelte";
	import TableHeaderRow from "./subcomponents/TableHeaderRow.svelte";
	import TableBody from "./subcomponents/TableBody.svelte";
	import type { ITableViewContext } from "./types";
	import { writable } from "svelte/store";
	import Loader from "../../common/loader/Loader.svelte";
	import PaginationControls from "../../common/paginationControls/PaginationControls.svelte";

	export let connection: connections.Connection;
	export let table: IDatabaseTable;

	let columns: queryBindingTypes.DatabaseColumn[] = undefined;
	let rows: { [key: string]: any }[] = undefined;
	let loading: boolean = true;
	let loadingData: boolean = true;
	let page: number = 0;
	let columnSizes: { [key: string]: number } = {};
	let details: queryBindingTypes.QueryResultTablePaginationData = undefined;

	const context = writable<ITableViewContext>({
		columns,
		connection,
		table,
		rows,
		columnSizes,
	});
	setContext("tableView", context);

	function handleExpandColumn(e: CustomEvent<{ name: string; width: number }>) {
		columnSizes[e.detail.name] = e.detail.width;
	}

	function updateContext(c: Partial<ITableViewContext>) {
		context.update(v => ({ ...v, ...c}));
	}
	$: updateContext({ columns, connection, table, rows, columnSizes });

	async function getTableData(c: connections.Connection, n: string) {
		loading = true;
		if (c && n) {
			columns = await GetTableData(c, n);
		}
		loading = false;
	}
	$: getTableData(connection, table.name);

	async function loadTableData(c: connections.Connection, tableName: string, page: number) {
		loadingData = true;
		if (c && tableName) {
			const result = (await LoadTableData(c, tableName, page));
			rows = result?.rows;
			details = result?.details;
		}
		loadingData = false;
	}
	$: loadTableData(connection, table.name, page);
</script>

{#if loading}
	<Loader />
{:else}
	<div class="container">
		{#if loadingData}
			<Loader />
		{:else}
			<div class="table-wrapper">
				<table class="table">
					<TableHeaderRow on:resize={handleExpandColumn} />
					<TableBody on:resize={handleExpandColumn} />
				</table>
			</div>
		{/if}
		<PaginationControls bind:offset={page} total={details?.total_pages} />
	</div>
{/if}

<style lang="scss">
	.container {
		height: 100%;
		display: grid;
		grid-template-rows: 1fr auto;
		gap: .5em;
		padding-bottom: .5em;
	}

	.table-wrapper {
		width: 100%;
		height: 100%;
		overflow: auto;
	}

	.table {
		height: fit-content;
		width: fit-content;
		border-collapse: separate;
		border-spacing: 0;
		background: var(--primary);
	}
</style>
