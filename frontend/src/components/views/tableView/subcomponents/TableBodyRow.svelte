<script lang="ts">
	import { getContext } from "svelte";
	import TableBodyCell from "./TableBodyCell.svelte";
	import type { ITableViewContext } from "../types";
	import TableCell from "./TableCell.svelte";
	import type { Writable } from "svelte/store";

	const context = getContext("tableView") as Writable<ITableViewContext>;
</script>

<tr>
	{#if Array.isArray($context?.columns)}
		{#each $context?.columns as column, index}
			<TableCell />
		{/each}
	{/if}
</tr>
