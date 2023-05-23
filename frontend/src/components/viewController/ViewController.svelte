<script lang="ts">
    import Tabs from "./subcomponents/Tabs.svelte";
</script>

<div class="view">
    <Tabs />
</div>

<style lang="scss">
    .view {
        grid-column: 2;
        grid-row: 3;
        height: 100%;
        width: 100%;
        overflow-y: auto;
        overflow-x: hidden;
    }
</style>
