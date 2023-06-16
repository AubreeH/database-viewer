<script lang="ts">
	import { getContext } from "svelte";
	import TableBodyCell from "./TableBodyCell.svelte";
	import type { ITableViewContext } from "../types";
	import TableCell from "./TableCell.svelte";
	import type { Writable } from "svelte/store";

	const context = getContext("tableView") as Writable<ITableViewContext>;

	export let row: { [key: string]: any };
</script>

<tr class="table-row">
	{#if row && Array.isArray($context.columns)}
		{#each $context.columns as column}
			<TableCell value={row[column.column.field]} />
		{/each}
	{/if}
</tr>

<style lang="scss">
	.table-row:hover {
		background: var(--hover);
	}
</style>
