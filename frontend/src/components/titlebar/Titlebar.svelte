<script lang="ts">

    import { handleTitlebarDoubleClick } from "./functions";

</script>

<div class="titlebar" on:dblclick={handleTitlebarDoubleClick}>

</div>

<style lang="scss">
    .titlebar {
        grid-column: 1/-1;
        grid-row: 1;
        height: 2em;
        border-bottom: 1px solid rgba(255, 255, 255, 0.1);
        --wails-draggable:drag;
    }
</style>