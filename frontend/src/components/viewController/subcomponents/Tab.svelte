<script lang="ts">
    import { closeTab, focusTab } from "../../../stores/mainView";
    import type { IDisplayComponent } from "../../../types/componentTypes";
    import ContextMenu from "../../common/contextMenu/ContextMenu.svelte";

    export let tab: IDisplayComponent<any>;
    export let index: number;
    export let focused: boolean;

    function handleClick() {
        focusTab(index);
    }

    function handleClose() {
        closeTab(index);
    }
</script>

<ContextMenu
    let:contextmenu
    id={`tab-${tab.name}-${index}`}
    options={[{ display: "Close", handler: handleClose }]}
>
    <button
        {...contextmenu}
        class={`tab primary hoverable${
            focused ? " highlighted" : " highlightable"
        }`}
        on:click={handleClick}
    >
        {tab.name}
    </button>
</ContextMenu>

<style lang="scss">
    .tab {
        width: fit-content;
        min-height: calc(2em - 1px);
        height: calc(2em - 1px);
        flex: 0 0 fit-content;
        border: 0;
        border-radius: 0;
        box-shadow: none;
        margin: 0;
        display: flex;
        align-items: center;
        justify-content: center;
    }
</style>
