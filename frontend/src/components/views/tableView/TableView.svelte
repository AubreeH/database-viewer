<script lang="ts">
	import { connections } from "../../../../wailsjs/go/models";
	import type { connectionsBinding } from "../../../../wailsjs/go/models";
	import type { IDatabaseTable } from "../../sidePanels/tableList/types";
	import { GetTableData } from "../../../../wailsjs/go/connectionsBinding/ConnectionsBinding";

	export let connection: connections.Connection;
	export let table: IDatabaseTable;

	let columns: connectionsBinding.DatabaseColumn = undefined;

	$: getTableData(connection, table.name);

	async function getTableData(c: connections.Connection, n: string) {
		if (c && n) {
			columns = await GetTableData(c, n);
		}
	}
</script>

<div class="column-row">
	{#if Array.isArray(columns)}
		{#each columns as column}
			<div class="column-header">
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
