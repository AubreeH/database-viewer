<script lang="ts">
	import { getContext } from "svelte";
	import type { queryBindingTypes } from "../../../../../wailsjs/go/models";
	import { openNewTableViewTab, openTableViewTab } from "../../../../stores/mainView";
	import type { ITableViewContext } from "../types";
	import type { Writable } from "svelte/store";

	export let column: queryBindingTypes.DatabaseColumn;

	const context = getContext("tableView") as Writable<ITableViewContext>;

	function handleDblClick(e: MouseEvent, column: queryBindingTypes.DatabaseColumn) {
		console.log(column);
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

<th class="table-header-cell" on:dblclick={(e) => handleDblClick(e, column)}>
	{column.column.field}
</th>

<style lang="scss">
	.table-header-cell {
		padding: 0.25em 2em;
		border-right: var(--border-outline-soft);
		font-weight: normal;
		user-select: none;
		-webkit-user-select: none;
	}
</style>
