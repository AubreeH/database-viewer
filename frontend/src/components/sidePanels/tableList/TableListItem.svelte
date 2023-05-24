<script lang="ts">
	import type { connections } from "../../../../wailsjs/go/models";
	import { openTabFocused } from "../../../stores/mainView";
	import TableView from "../../views/tableView/TableView.svelte";
	import type { IDatabaseTable } from "./types";

	export let table: IDatabaseTable;
	export let connection: connections.Connection;

	function handleClick() {
		openTabFocused({
			component: TableView,
			props: { table, connection },
			name: `Table: ${table.name}`,
		});
	}
</script>

<button class="primary hoverable" on:click={handleClick}>
	{table.name}
</button>
