<script lang="ts">
	import { GetTableData, LoadTableData } from "../../../../wailsjs/go/queryBinding/QueryBinding";
	import type { connections, queryBindingTypes } from "../../../../wailsjs/go/models";
	import type { IDatabaseTable } from "../../sidePanels/tableList/types";
	import Loader from "../../common/loader/Loader.svelte";
	import PaginationControls from "../../common/paginationControls/PaginationControls.svelte";

	export let connection: connections.Connection;
	export let table: IDatabaseTable;

	let columns: queryBindingTypes.DatabaseColumn[] = undefined;
	let rows: { [key: string]: any }[] = undefined;
	let loading: boolean = true;
	let page: number = 0;
	let columnSizes: { [key: string]: number } = {};

	let dragStartClientXPosition = undefined;
	let dragStartWidth = undefined;

	let hoveringColumnResizer: {columnName: string, rowIndex: number} | undefined = undefined;
	let draggingColumnResizer: {columnName: string, rowIndex: number} | undefined = undefined;

	async function getTableData(c: connections.Connection, n: string) {
		loading = true;
		if (c && n) {
			columns = await GetTableData(c, n);
		}
		loading = false;
	}
	$: getTableData(connection, table.name);

	async function loadTableData(c: connections.Connection, tableName: string, page: number) {
		if (c && tableName) {
			rows = (await LoadTableData(c, tableName, page))?.rows;
		}
	}
	$: loadTableData(connection, table.name, page);

	function getColumnWidth(columnName: string) {
		if (columnSizes[columnName]) {
			return columnSizes[columnName];
		}
		
		const header = document.getElementById(`table-header-cell_${columnName}`);
		if (header) {
			return header.clientWidth;
		}

		return 100;
	}

	function handleBeginHoverColumnResizer(e: Event, columnName: string, rowIndex: number) {
		hoveringColumnResizer = {columnName, rowIndex};
	}

	function handleEndHoverColumnResizer(e: Event, columnName: string, rowIndex: number) {
		hoveringColumnResizer = undefined;
	}

	function handleKeyDownColumnResizer(e: KeyboardEvent, columnName: string, rowIndex: number) {
		let columnSize = getColumnWidth(columnName);

		if (e.key === "ArrowLeft") {
			columnSize -= 1;
		} else if (e.key === "ArrowRight") {
			columnSize += 1;
		}

		columnSizes[columnName] = columnSize;
	}

	function handleDragColumnResizer(e: MouseEvent, columnName: string, rowIndex: number) {
		if (e.clientX && dragStartClientXPosition !== undefined && dragStartWidth !== undefined) {
			const movement = e.clientX - dragStartClientXPosition;
			columnSizes[columnName] = dragStartWidth + movement;
		}
	}

	function handleDragStartColumnResizer(e: MouseEvent, columnName: string, rowIndex: number) {
		function handleDrag(e: MouseEvent) {
			handleDragColumnResizer(e, columnName, rowIndex);
		}

		function handleDragEnd(e: MouseEvent) {
			handleDragEndColumnResizer(e, columnName, rowIndex);
			window.removeEventListener("mouseup", handleDragEnd);
			window.removeEventListener("mousemove", handleDrag);
		}

		window.addEventListener("mouseup", handleDragEnd)
		window.addEventListener("mousemove", handleDrag)
		dragStartClientXPosition = e.clientX;
		dragStartWidth = getColumnWidth(columnName);
		draggingColumnResizer = {columnName, rowIndex};
	}

	function handleDragEndColumnResizer(e: MouseEvent, columnName: string, rowIndex: number) {
		dragStartClientXPosition = undefined;
		dragStartWidth = undefined;
		draggingColumnResizer = undefined;
		window.removeEventListener("mouseup", (ev: MouseEvent) => handleDragEndColumnResizer(ev, columnName, rowIndex))
	}
</script>

{#if loading}
	<Loader />
{:else}
	<div class="container">
		<div class="table-wrapper">
			<table class="table">
				<thead>
					<tr class={'table-header-row'}>
						{#each columns as column}
							<th
								id={`table-header-cell_${column.column.field}`}
								class="table-header-cell" 
								style={columnSizes[column.column.field] ? `width: ${columnSizes[column.column.field]}px` : undefined}
							>
								{column.column.field}
							</th>
							<td class={`column-divider`}>
								<button
									on:keydown={(e) => handleKeyDownColumnResizer(e, column.column.field, -1)}
									on:mousedown={(e) => handleDragStartColumnResizer(e, column.column.field, -1)}
									on:mouseenter={(e) => handleBeginHoverColumnResizer(e, column.column.field, -1)}
									on:focus={(e) => handleBeginHoverColumnResizer(e, column.column.field, -1)}
									on:mouseleave={(e) => handleEndHoverColumnResizer(e, column.column.field, -1)}
									on:blur={(e) => handleEndHoverColumnResizer(e, column.column.field, -1)}
								/>
							</td>
						{/each}
					</tr>
				</thead>
				<tbody>
					{#each rows as row, i}
						<tr class={`table-body-row${!!i ? " show-border" : ""}`}>
							{#each columns as column}
								<td class="table-cell">
									{row[column.column.field]}
								</td>
								<td class={`column-divider${hoveringColumnResizer?.columnName === column.column.field ? " hovering" : ""}${draggingColumnResizer?.columnName === column.column.field ? " dragging" : ""}`}>
									<button
										on:keydown={(e) => handleKeyDownColumnResizer(e, column.column.field, i)}
										on:mousedown={(e) => handleDragStartColumnResizer(e, column.column.field, -1)}
										on:mouseenter={(e) => handleBeginHoverColumnResizer(e, column.column.field, i)}
										on:focus={(e) => handleBeginHoverColumnResizer(e, column.column.field, i)}
										on:mouseleave={(e) => handleEndHoverColumnResizer(e, column.column.field, i)}
										on:blur={(e) => handleEndHoverColumnResizer(e, column.column.field, i)}
									/>
								</td>
							{/each}
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
		<PaginationControls bind:offset={page} total={10} />
	</div>
{/if}

<style lang="scss">
	.container {
		height: 100%;
		display: grid;
		grid-template-rows: 1fr auto;
		gap: .5em;
		padding-bottom: .5em;
	}

	.table-wrapper {
		width: 100%;
		height: 100%;
		overflow: auto;
	}

	.table {
		height: fit-content;
		width: fit-content;
		border-collapse: separate;
		border-spacing: 0;
		background: var(--primary);
	}

	.table-header-row {
		background: var(--primary-contrast);
	}

	.table-header-cell {
		min-width: 5em;
		overflow-x: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		padding: 0.25em 2em;
		font-weight: normal;
		user-select: none;
		-webkit-user-select: none;
		box-sizing: border-box;
	}

	.table-body-row {
		&.show-border {
			--top-border: var(--border-outline-soft);
		}

		&:hover {
			background: var(--hover);			
		}
	}

	.table-cell {
		padding: 0.25em 1em;
		white-space: nowrap;
		overflow-x: auto;

		user-select: none;

		// cursor: text;
		
		border-top: var(--top-border);

		&:hover {
			background: var(--hover);
		}
	}

	.column-divider {
		padding: 0;
			
		border-top: var(--top-border);
		
		button {
			appearance: none;
			background: transparent;
			border: none;
			outline: none;
			box-shadow: none;
			padding: 0;
			height: 100%;
			width: 5px;
			margin-inline: -2px;

			cursor: ew-resize;
			
			display: grid;
			place-items: center;
			position: relative;

			&::after {
				content: "";
				width: 1px;
				height: 100%;
				display: flex;
				background: var(--outline-very-soft);
				pointer-events: none;
				transition: width 10ms ease;
			}
		}

		&.hovering button::after {
			// background: var(--outline-soft);
			width: 2px;
		}

		&.dragging button {
			width: 21px;
			margin-inline: -10px;
		}
	}
</style>
