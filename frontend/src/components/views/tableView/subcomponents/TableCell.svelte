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
		if (active) {
			dispatch("resize", { name: column?.column?.field, width: columnWidth });
		}
	}

	let cellInner: HTMLElement = undefined;
	$: observer = new ResizeObserver(function(mutations) {
		mutations.forEach(function(mutation) {
			columnWidth = mutation.contentRect.width;
		});
	});

	$: {
		if (cellInner) {
			observer.observe(cellInner);
		}
	}

	function handleMouseDown() {
		function handleResetActive() {
			active = false;
			window.removeEventListener('mouseup', handleResetActive)
		}

		active = true;
		window.addEventListener('mouseup', handleResetActive)
	}
</script>

<td class="table-cell">
	<div 
		class="table-cell__inner" 
		style={!active ? `width: ${width} !important;` : undefined} 
		bind:this={cellInner}
		on:mousedown={handleMouseDown}
	>
		<div class="text-wrapper">
			{value}
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
		padding: 0.25em 1em;
	}
</style>
