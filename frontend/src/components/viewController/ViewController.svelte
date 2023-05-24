<script lang="ts">
	import Tabs from "./subcomponents/Tabs.svelte";
	import { MainViewStore } from "../../stores/mainView";

	function getFocusedTab(activeTab: number) {
		return $MainViewStore.tabs[activeTab];
	}

	$: activeTab = $MainViewStore.activeTab;
	$: component = !isNaN(Number(activeTab))
		? $MainViewStore?.tabs[activeTab]?.component
		: null;
	$: props = !isNaN(Number(activeTab))
		? $MainViewStore?.tabs[activeTab]?.props
		: null;

	// $: {
	// 	console.log(activeTab, component, props);
	// }
</script>

<div class="view">
	<Tabs />
	<div class="main-view">
		{#if $MainViewStore.activeTab !== undefined && $MainViewStore.activeTab > -1 && getFocusedTab($MainViewStore.activeTab)}
			<svelte:component this={component} {...props} />
		{/if}
	</div>
</div>

<style lang="scss">
	.view {
		grid-column: 2;
		grid-row: 3;
		height: 100%;
		width: 100%;
		overflow-y: auto;
		overflow-x: hidden;

		display: grid;
		grid-template-rows: auto 1fr;
	}

	.main-view {
		border-top: 1px solid var(--divider);
	}
</style>
