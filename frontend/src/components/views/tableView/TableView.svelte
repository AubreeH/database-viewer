<script lang="ts">
	import { GetTableData } from "../../../../wailsjs/go/queryBinding/QueryBinding";
	import type { connections, queryBindingTypes } from "../../../../wailsjs/go/models";
	import type { IDatabaseTable } from "../../sidePanels/tableList/types";
	import { openNewTableViewTab, openTableViewTab } from "../../../stores/mainView";

	export let connection: connections.Connection;
	export let table: IDatabaseTable;

	let columns: queryBinding.DatabaseColumn = undefined;

	$: getTableData(connection, table.name);

	async function getTableData(c: connections.Connection, n: string) {
		if (c && n) {
			columns = await GetTableData(c, n);
			console.log(columns);
		}
	}

	function handleDblClick(e: MouseEvent, column: queryBinding.DatabaseColumn) {
		if (Array.isArray(column?.Indexes) && column.Indexes.length) {
			const foreignKeyIndex = column.Indexes.find((i) => i?.RefColumn?.String !== "");
			if (foreignKeyIndex !== -1) {
				if (e.ctrlKey) {
					openNewTableViewTab(connection, column?.ForeignKey?.ReferencedTableName?.String);
				} else {
					openTableViewTab(connection, column?.ForeignKey?.ReferencedTableName?.String);
				}
			}
		}
	}
</script>

<div class="column-row">
	{#if Array.isArray(columns)}
		{#each columns as column}
			<div class="column-header" on:dblclick={(e) => handleDblClick(e, column)}>
				{column.Column.Field}
			</div>
		{/each}
	{/if}
</div>

<style lang="scss">
	.column-row {
		display: flex;
		flex-direction: row;
		background: var(--primary-contrast);
	}

	.column-header {
		padding: 0.25em 2em;
		border-right: var(--border-outline-soft);
	}
</style>
