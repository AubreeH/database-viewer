<script lang="ts">
	import { getContext } from "svelte";
	import type { queryBindingTypes } from "../../../../../wailsjs/go/models";
	import { openNewTableViewTab, openTableViewTab } from "../../../../stores/mainView";
	import type { ITableViewContext } from "../types";
	import type { Writable } from "svelte/store";

	export let column: queryBindingTypes.DatabaseColumn;

	const context = getContext("tableView") as Writable<ITableViewContext>;

	function handleDblClick(e: MouseEvent, column: queryBindingTypes.DatabaseColumn) {
		if (Array.isArray(column?.Indexes) && column.Indexes.length) {
			const foreignKeyIndex = column.Indexes.find((i) => i?.RefColumn?.String !== "");
			if (foreignKeyIndex !== -1) {
				if (e.ctrlKey) {
					openNewTableViewTab($context.connection, column?.ForeignKey?.ReferencedTableName?.String);
				} else {
					openTableViewTab($context.connection, column?.ForeignKey?.ReferencedTableName?.String);
				}
			}
		}
	}
</script>

<th class="table-header-cell" on:dblclick={(e) => handleDblClick(e, column)}>
	{column.Column.Field}
</th>

<style lang="scss">
	.table-header-cell {
		padding: 0.25em 2em;
		border-right: var(--border-outline-soft);
		font-weight: normal;
	}
</style>
