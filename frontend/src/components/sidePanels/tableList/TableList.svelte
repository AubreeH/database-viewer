<script lang="ts">
	import type { connections } from "../../../../wailsjs/go/models";
	import TableListItem from "./TableListItem.svelte";
	import type { IDatabaseTable } from "./types";

	export let tables: IDatabaseTable[] = undefined;
	export let connection: connections.Connection;
</script>

{#if Array.isArray(tables)}
	{#each tables as table}
		<TableListItem {table} {connection} />
	{/each}
{/if}
