<script lang="ts">
	import ContextMenu from "../../common/contextMenu/ContextMenu.svelte";
	import type { connections } from "../../../../wailsjs/go/models";
	import {
		openNewTableViewTab,
		openTableViewTab,
	} from "../../../stores/mainView";
	import type { IDatabaseTable } from "./types";
	import Icon from "../../common/icon/Icon.svelte";

	export let table: IDatabaseTable;
	export let connection: connections.Connection;

	function handleClick() {
		openTableViewTab(connection, table.name);
	}

	function handleOpenNew() {
		openNewTableViewTab(connection, table.name);
	}
</script>

<ContextMenu
	id={`table-list-item-${connection.name}-${table.name}`}
	let:contextmenu
	options={[{ display: "Open New", handler: handleOpenNew }]}
>
	<button {...contextmenu} class="primary hoverable table-list-item" on:click={handleClick}>
		<Icon class="table-list-item-chevron">
			chevron_right
		</Icon>
		
		<span title={table.name}>
			{table.name}
		</span>
	</button>
</ContextMenu>

<style lang="scss">
	.table-list-item {
		display: grid;
		grid-template-columns: auto 1fr;
		justify-items: start;
		align-items: center;
		gap: .5em;
		line-height: 1;
		vertical-align: middle;
		padding: .375em .625em;
		margin-right: 1em;

		:global(.table-list-item-chevron) {
			height: 1.75em;
			width: 1.75em;
			display: grid;
			align-content: center;
			border-radius: .25em;
			&:hover {
				box-shadow: 0 0 .5em var(--box-shadow-light) inset;
			}
		}

		span {
			width: 100%;

			overflow: hidden;
			white-space: nowrap;
			text-overflow: ellipsis;
			text-align: start;
		}
	}

	
</style>
