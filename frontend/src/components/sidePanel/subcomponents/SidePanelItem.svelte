<script lang="ts">
    import Icon from "../../common/icon/Icon.svelte";

    export let name = "";
    export let details = "";
    export let connected = false;
</script>

<div class="side-panel-item hoverable">
    <h4>{name}</h4>
    {#if connected}
        <Icon>check_circle</Icon>
    {/if}
    <p>{details}</p>
</div>

<style lang="scss">
    .side-panel-item {
        padding: .8em;
        background-color: #404048;
        border-radius: .5em;
        display: grid;
        grid-template-columns: 1fr auto;

        & > h4 {
            text-align: left;
            margin: 0;
            text-transform: capitalize;
            font-size: .9em;

            grid-column: 1;
            grid-row: 1;
        }

        & > :global(.icon) {
            font-size: 1em;
            color: #20ff50;
        }

        & > p {
            grid-column: 1;
            grid-row: 2;
        }
    }
</style>