<script>
    import SidePanelItem from "./subcomponents/SidePanelItem.svelte";
</script>

<div class="side-panel">
    <SidePanelItem name="test1" />
    <SidePanelItem name="test2" connected />
    <SidePanelItem name="test3" />
    <SidePanelItem name="test4" connected />
</div>

<style lang="scss">
    .side-panel {
        grid-column: 1;
        grid-row: 3;
        height: 100%;
        width: 100%;
        border-right: 1px solid rgba(255, 255, 255, 0.1);
        display: flex;
        flex-direction: column;
        padding: 1em;
        gap: 1em;
    }
</style>