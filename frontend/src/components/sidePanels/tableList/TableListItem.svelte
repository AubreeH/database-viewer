<script lang="ts">
	import ContextMenu from "../../common/contextMenu/ContextMenu.svelte";
	import type { connections } from "../../../../wailsjs/go/models";
	import {
		openNewTableViewTab,
		openTableViewTab,
	} from "../../../stores/mainView";
	import type { IDatabaseTable } from "./types";

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
	<button {...contextmenu} class="primary hoverable" on:click={handleClick}>
		{table.name}
	</button>
</ContextMenu>
