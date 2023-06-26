<script lang="ts">
	import type { connections } from "../../../../wailsjs/go/models";
	import { GetTables } from "../../../../wailsjs/go/queryBinding/QueryBinding";
	import TableListItem from "./TableListItem.svelte";
	import type { IDatabaseTable } from "./types";

	export let connection: connections.Connection;
	export let connectionName: string = undefined
	let tables: IDatabaseTable[] = undefined
	let offset = 0;
	let loading = true;

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

		const tableNameArray = await GetTables(connection, "")
		tables = Array.isArray(tableNameArray) ? tableNameArray.map((t) => ({ name: t })) : []
		
		loading = false;
	}

</script>

{#if Array.isArray(tables)}
	{#each tables as table}
		<TableListItem {table} {connection} />
	{/each}
{/if}
