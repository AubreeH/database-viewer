<script lang="ts">
	import { getContext } from "svelte";
	import TableBodyRow from "./TableBodyRow.svelte";
	import type { ITableViewContext } from "../types";
	import type { Writable } from "svelte/store";

	const context = getContext("tableView") as Writable<ITableViewContext>;
	$: console.log($context?.rows);
</script>

<tbody>
	{#if Array.isArray($context?.rows)}
		{#each $context?.rows as row}
			<TableBodyRow {row} />
		{/each}
	{/if}
</tbody>
