<script lang="ts">
    import { openNewSidePanel } from "../../../stores/sidePanel";
    import Icon from "../../common/icon/Icon.svelte";
    import Test from "./Test.svelte";
    import { DeleteConnection } from "../../../../wailsjs/go/connectionsbinding/ConnectionsBinding";

    export let name = "";
    export let details = "";
    export let connected = false;

    function handleClick(e: MouseEvent) {
        handleDeleteConnection();
    }

    function handleDeleteConnection() {
        DeleteConnection(name);
    }
</script>

<button
    class="side-panel-item hoverable primary"
    on:click={handleClick}
    on:dblclick={handleDeleteConnection}
>
    <h4>{name}</h4>
    {#if connected}
        <Icon name="check_circle" />
    {/if}
    <p>{details}</p>
</button>

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
