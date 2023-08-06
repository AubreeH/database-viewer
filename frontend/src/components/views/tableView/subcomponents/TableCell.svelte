<script lang="ts">
	import { createEventDispatcher, getContext } from "svelte";
	import type { Writable } from "svelte/store";
	import type { ITableViewContext } from "../types";
	import type { queryBindingTypes } from "../../../../../wailsjs/go/models";
	import { getKeybindManagerContext } from "../../../common/keybindManager/functions";

	export let value: any = undefined;
	export let column: queryBindingTypes.DatabaseColumn = undefined;
	const context = getContext("tableView") as Writable<ITableViewContext>;
	let hovering = false;
	let active = false;

	$: width = $context.columnSizes[column.column.field] ? `${$context.columnSizes[column.column.field]}px` : "100%"

	const dispatch = createEventDispatcher();

	let columnWidth = 0;
	$: handleUpdateColumnWidth(columnWidth);
	function handleUpdateColumnWidth(columnWidth: number) {
		if (hovering) {
			dispatch("resize", { name: column?.column?.field, width: columnWidth });
		}
	}

	let cellInner: HTMLElement = undefined;
</script>

<td class="table-cell">
	<div 
		class="table-cell__inner" 
		style={!hovering ? `width: ${width} !important;` : undefined} 
		bind:this={cellInner}
	>
		<div class={`text-wrapper${value===null?" null-value":""}${value===""?" empty-value":""}`}>
			{value === null ? "null" : value === "" ? "empty" : value}
		</div>
	</div>
</td>

<style lang="scss">
	.table-cell {
		padding: 0;
		white-space: nowrap;
		overflow-x: auto;
		
		border-top: var(--top-border);
		&:not(:last-child) {
			border-right: var(--border-outline-very-soft);
		}

		&:hover {
			background: var(--hover);
		}
	}
	.table-cell__inner {
		height: 100%;
  		overflow-x: auto;
		overflow-y: hidden;

		box-sizing: border-box;
		overflow-x: hidden;
		min-width: 5em;

		display: flex;

		&:active {
			background: red;
		}
	}

	.table-cell__inner:hover, .table-cell__inner:active {
		resize: horizontal;
	}
	.table-cell__inner:not(:hover, :active) {
		width: 100% !important;
		resize: none;
	}

	.table-cell__inner::-webkit-resizer {
		appearance: none;
	}
	.text-wrapper {
		margin: 0.25em 1em;

		&.null-value,
		&.empty-value {
			opacity: .5;
		}
	}
</style>
