<script lang="ts">
	import { onMount } from "svelte";
	import SidePanelItem from "./subcomponents/SidePanelItem.svelte";
	import SidePanelStore, { goToPreviewSidePanel } from "../../stores/sidePanel";
	import Icon from "../common/icon/Icon.svelte";
	import { connections } from "../../../wailsjs/go/models";
	import { LoadConnections } from "../../../wailsjs/go/connectionsBinding/ConnectionsBinding";
	import { openModal } from "../../stores/modal";
	import ModifyConnectionModal from "../modals/newConnectionModal/ModifyConnectionModal.svelte";
	import { EventsOn } from "../../../wailsjs/runtime/runtime";
	import { getNotificationContext } from "../notifications/functions";

	let loading: boolean = true;
	let connectionsList: connections.Connection[] = undefined;

	const { sendFatalNotification } = getNotificationContext();

	function handleOpenCreateNewConnectionModal() {
		openModal({
			title: "Create New Connection",
			component: ModifyConnectionModal,
		});
	}

	onMount(async () => {
		try {
			connectionsList = await LoadConnections();
			loading = false;
			EventsOn("connections-updated", (data) => {
				if (Array.isArray(data)) {
					connectionsList = data.map((c) => new connections.Connection(c));
				} else {
					connectionsList = undefined;
				}
			});
		} catch (e) {
			sendFatalNotification({
				message: "",
			});
		}
	});
</script>

<div class="side-panel">
	{#if Array.isArray($SidePanelStore) && $SidePanelStore.length}
		<div class="side-panel-header">
			<button class="back primary" on:click={goToPreviewSidePanel}>
				<Icon>chevron_left</Icon>
			</button>
			<h4>
				{$SidePanelStore.slice(-1)[0].name ?? ""}
			</h4>
		</div>
	{/if}
	<div class="side-panel-main hidden-scrollbar">
		{#if Array.isArray($SidePanelStore) && $SidePanelStore.length}
			<svelte:component this={$SidePanelStore.slice(-1)[0].component} {...$SidePanelStore.slice(-1)[0]?.props ?? []} />
		{:else if Array.isArray(connectionsList)}
			{#each connectionsList as connection}
				<SidePanelItem {connection} />
			{/each}
		{/if}
	</div>
	{#if !Array.isArray($SidePanelStore) || !$SidePanelStore.length}
		<div class="side-panel-footer">
			<button class="small icon hoverable primary" title="New Connection" on:click={handleOpenCreateNewConnectionModal}> add </button>
		</div>
	{/if}
</div>

<style lang="scss">
	.side-panel {
		max-height: 100%;
		height: 100%;
		overflow-y: auto;

		width: 100%;
		padding: 0;

		grid-column: 1;
		grid-row: 3;

		display: grid;
		grid-template-rows: auto 1fr auto;
		gap: 1em;

		border-right: 1px solid var(--divider);
	}

	.side-panel-header {
		grid-row: 1;

		padding: 1em 1em 0;

		display: grid;
		grid-template-columns: 1fr 3fr 1fr;
		align-items: center;
	}

	.side-panel-main {
		grid-row: 2;

		display: flex;
		flex-direction: column;
		gap: 1em;

		overflow-y: auto;
		padding: 0;
	}

	.side-panel-footer {
		grid-row: 3;

		display: flex;
		align-items: center;
		justify-content: end;

		button {
			width: fit-content;
		}
	}

	.back {
		width: fit-content;
	}
</style>
