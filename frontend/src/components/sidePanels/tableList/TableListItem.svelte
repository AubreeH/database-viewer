<script lang="ts">
	import ContextMenu from "../../common/contextMenu/ContextMenu.svelte";
	import type { connections } from "../../../../wailsjs/go/models";
	import { openNewTabFocused, openTabFocused } from "../../../stores/mainView";
	import TableView from "../../views/tableView/TableView.svelte";
	import type { IDatabaseTable } from "./types";

	export let table: IDatabaseTable;
	export let connection: connections.Connection;

	function handleClick() {
		openTabFocused({
			component: TableView,
			props: { table, connection },
			name: `Table: ${table.name}`,
			id: table.name,
		});
	}

	function handleOpenNew() {
		openNewTabFocused({
			component: TableView,
			props: { table, connection },
			name: `Table: ${table.name}`,
			id: table.name,
		});
	}
</script>

<ContextMenu
	id={`table-list-item-${connection.name}-${table.name}`}
	let:contextmenu
	options={[{ display: "Open New", handler: handleOpenNew }]}
>
	<button {...contextmenu} class="primary hoverable" on:click={handleClick}>
		{table.name}
	</button>
</ContextMenu>
