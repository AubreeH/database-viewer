<script lang="ts">
	import { getContext } from "svelte";
	import type { ITableViewContext } from "../types";
	import TableCell from "./TableCell.svelte";
	import type { Writable } from "svelte/store";

	const context = getContext("tableView") as Writable<ITableViewContext>;

	export let showBorder: boolean = true;

	export let row: { [key: string]: any };
</script>

<tr class={`table-row${showBorder ? " show-border" : ""}`}>
	{#if row && Array.isArray($context.columns)}
		{#each $context.columns as column}
			<TableCell {column} on:resize value={row[column.column.field]} />
		{/each}
	{/if}
</tr>

<style lang="scss">
</style>
