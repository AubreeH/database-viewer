<script lang="ts">
	import { getContext } from "svelte";
	import TableHeaderCell from "./TableHeaderCell.svelte";
	import type { Writable } from "svelte/store";
	import type { ITableViewContext } from "../types";

	const context = getContext("tableView") as Writable<ITableViewContext>;
</script>

<thead class="table-header-row">
	<tr>
		{#if Array.isArray($context?.columns)}
			{#each $context.columns as column}
				<TableHeaderCell on:resize {column} />
			{/each}
		{/if}
	</tr>
</thead>

<style lang="scss">
	.table-header-row {
		// position: sticky;
		// top: 0;
	}
</style>
