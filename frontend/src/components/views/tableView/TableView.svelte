<script lang="ts">
	import { GetTableData } from "../../../../wailsjs/go/queryBinding/QueryBinding";
	import type {
		connections,
		queryBinding,
	} from "../../../../wailsjs/go/models";
	import type { IDatabaseTable } from "../../sidePanels/tableList/types";

	export let connection: connections.Connection;
	export let table: IDatabaseTable;

	let columns: queryBinding.DatabaseColumn = undefined;

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
