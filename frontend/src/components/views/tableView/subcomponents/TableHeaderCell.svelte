<script lang="ts">
	import { createEventDispatcher, getContext } from "svelte";
	import type { queryBindingTypes } from "../../../../../wailsjs/go/models";
	import { openNewTableViewTab, openTableViewTab } from "../../../../stores/mainView";
	import type { ITableViewContext } from "../types";
	import type { Writable } from "svelte/store";

	export let column: queryBindingTypes.DatabaseColumn;
	const context = getContext("tableView") as Writable<ITableViewContext>;

	const dispatch = createEventDispatcher();

	$: columnWidth = $context.columnSizes[column.column.field] || 100;

	let cellInner: HTMLElement = undefined;
	// $: observer = new ResizeObserver(function(mutations) {
	// 	mutations.forEach(function(mutation) {
	// 		columnWidth = mutation.contentRect.width;
	// 	});
	// });

	// $: {
	// 	if (cellInner) {
	// 		observer.observe(cellInner);
	// 	}
	// }

	function handleDblClick(e: MouseEvent, column: queryBindingTypes.DatabaseColumn) {
		if (Array.isArray(column?.indexes) && column.indexes.length) {
			const foreignKeyIndex = column.indexes.findIndex((i) => i?.ref_column?.String !== "");
			if (foreignKeyIndex !== -1) {
				if (e.ctrlKey) {
					openNewTableViewTab($context.connection, column?.indexes[foreignKeyIndex]?.ref_table?.String);
				} else {
					openTableViewTab($context.connection, column?.indexes[foreignKeyIndex]?.ref_table?.String);
				}
			}
		}
	}
</script>

<th class="table-header-cell" style="width: {columnWidth}px;" on:dblclick={(e) => handleDblClick(e, column)} bind:this={cellInner}>
	{column.column.field}
</th>

<style lang="scss">
	.table-header-cell {
		min-width: 5em;
		overflow-x: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		padding: 0.25em 2em;
		border-right: var(--border-outline-soft);
		font-weight: normal;
		user-select: none;
		-webkit-user-select: none;
		box-sizing: border-box;
	}
</style>
