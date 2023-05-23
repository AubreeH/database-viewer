<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import {
        MainViewStore,
        closeFocusedTab,
        focusTab,
    } from "../../../stores/mainView";
    import Tab from "./Tab.svelte";
    import { getKeybindManagerContext } from "../../common/keybindManager/functions";

    const { subscribe, unsubscribe } = getKeybindManagerContext();

    function onTabsKeyDown(e: KeyboardEvent) {
        if (e.key === "ArrowRight" || e.key === "ArrowLeft") {
            handleArrowKey(e);
        }
    }

    function handleArrowKey(e: KeyboardEvent) {
        e.preventDefault();
        e.stopPropagation();
        switch (e.key) {
            case "ArrowRight":
                focusTab($MainViewStore.activeTab + 1);
                break;
            case "ArrowLeft":
                focusTab($MainViewStore.activeTab - 1);
                break;
        }
    }

    onMount(() => {
        subscribe({
            ctrl: true,
            key: "w",
            name: "close tab",
            handler: closeFocusedTab,
        });
    });

    onDestroy(() => {
        unsubscribe({
            ctrl: true,
            key: "w",
            name: "close tab",
            handler: closeFocusedTab,
        });
    });
</script>

<div class="tabs hidden-scrollbar" on:keydown={onTabsKeyDown}>
    {#if Array.isArray($MainViewStore?.tabs)}
        {#each $MainViewStore?.tabs as tab, index}
            <Tab {tab} {index} focused={$MainViewStore.activeTab === index} />
        {/each}
    {/if}
</div>

<style lang="scss">
    .tabs {
        width: 100%;
        height: 2em;
        display: flex;
        gap: 3px;
        border-bottom: 1px solid rgba(255, 255, 255, 0.1);
        overflow-x: auto;
        background: var(--primary-contrast);
    }
</style>
