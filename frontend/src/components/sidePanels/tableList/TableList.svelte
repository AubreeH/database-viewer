<script lang="ts">
	import type { connections, queryBindingTypes } from "../../../../wailsjs/go/models";
	import { GetTables } from "../../../../wailsjs/go/queryBinding/QueryBinding";
	import Icon from "../../common/icon/Icon.svelte";
	import PaginationControls from "../../common/paginationControls/PaginationControls.svelte";
	import TableListItem from "./TableListItem.svelte";
	import type { IDatabaseTable } from "./types";

	export let connection: connections.Connection;
	export let connectionName: string = undefined
	let tables: IDatabaseTable[] = undefined
	let offset = 0;
	let loading = true;
	let details: queryBindingTypes.GetPaginationDetailsResult = undefined;

	$: handleConnectionUpdate(connection)
	async function handleConnectionUpdate(connection: connections.Connection) {
		if (connectionName === undefined) {
			connectionName = connection.name
		} else {
			if (connection.name !== connection.name) {
				offset = 0
				connectionName = connection.name
			}
		}
	}

	$: getTables(offset)
	async function getTables(offset: number) {
		loading = true;
		const result = await GetTables(connection, "", offset)
		tables = Array.isArray(result.tables) ? result.tables.map((t) => ({ name: t })) : []
		details = result.details		
		loading = false;
	}

</script>

<div class="table-list">
	<div class="table-list-items">
		{#if Array.isArray(tables)}
			{#each tables as table}
				<TableListItem {table} {connection} />
			{/each}
		{/if}
	</div>
	<PaginationControls bind:offset={offset} total={details?.total_pages} />
</div>



<style lang="scss">
	.table-list {
		height: 100%;
		display: grid;
		gap: 1em;
		overflow-y: auto;
		grid-template-rows: 1fr auto;
	}

	.table-list-items {
		display: grid;
		gap: .5em;
		overflow-y: auto;
		align-content: start;
		padding-left: 1em;
	}
</style>