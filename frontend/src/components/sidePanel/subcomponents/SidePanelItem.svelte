<script lang="ts">
	import { openNewSidePanel } from "../../../stores/sidePanel";
	import Icon from "../../common/icon/Icon.svelte";
	import {
		CloseConnection,
		DeleteConnection,
		GetTables,
		OpenConnection,
	} from "../../../../wailsjs/go/connectionsBinding/ConnectionsBinding";
	import ContextMenu from "../../common/contextMenu/ContextMenu.svelte";
	import type { connections } from "../../../../wailsjs/go/models";
	import TableList from "../../sidePanels/tableList/TableList.svelte";
	import { openModal } from "../../../stores/modal";
	import ModifyConnectionModal from "../../../components/modals/newConnectionModal/ModifyConnectionModal.svelte";

	export let connection: connections.Connection = undefined;

	async function handleClick(e: MouseEvent) {
		openNewSidePanel({
			component: TableList,
			name: `Database: ${connection.name}`,
			props: {
				connection,
				tables: (await GetTables(connection, "")).map((t) => ({ name: t })),
			},
		});
	}

	function handleDeleteConnection() {
		DeleteConnection(connection);
	}

	function handleConnect() {
		OpenConnection(connection);
	}

	function handleDisconnection() {
		CloseConnection(connection);
	}

	function handleEditConnection() {
		openModal({
			title: "Edit Connection",
			component: ModifyConnectionModal,
			props: { connection },
		});
	}
</script>

<ContextMenu
	id={`side-panel-item-${connection.name}`}
	let:contextmenu
	options={[
		connection.connected
			? { display: "Disconnect", handler: handleDisconnection }
			: { display: "Connect", handler: handleConnect },
		{ display: "Delete", handler: handleDeleteConnection },
		{ display: "Edit", handler: handleEditConnection },
	]}
>
	<button
		{...contextmenu}
		class="side-panel-item hoverable primary"
		on:click={handleClick}
		on:dblclick={handleDeleteConnection}
	>
		<h4>{connection.name}</h4>
		{#if connection.connected}
			<Icon name="check_circle" />
		{/if}
		<p>{connection.database}</p>
	</button>
</ContextMenu>

<style lang="scss">
	.side-panel-item {
		padding: 0.8em;
		background-color: var(--primary);
		display: grid;
		grid-template-columns: 1fr auto;

		& > h4 {
			text-align: left;
			margin: 0;
			text-transform: capitalize;

			grid-column: 1;
			grid-row: 1;
		}

		& > :global(.icon) {
			color: var(--active);
		}

		& > p {
			grid-column: 1;
			grid-row: 2;
		}
	}
</style>
