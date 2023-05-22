<script lang="ts">
    import { SaveConnection } from "../../../../wailsjs/go/connectionsbinding/ConnectionsBinding";
    import { connections } from "../../../../wailsjs/go/models";
    import { WindowReload } from "../../../../wailsjs/runtime/runtime";
    import { closeModal } from "../../../stores/modal";

    let name: string = "";
    let host: string = "";
    let port: string = "";
    let user: string = "";
    let database: string = "";

    let loading = false;

    async function handleCreateNewConnection() {
        loading = true;

        await SaveConnection(
            new connections.Connection({
                name,
                host,
                port,
                user,
                database,
            })
        );

        WindowReload();
        closeModal();
    }
</script>

<div class="create-new-connection-form">
    <div class="field">
        <label for="name-input">Connection Name: </label>
        <input id="name-input" bind:value={name} />
    </div>

    <div class="field">
        <label for="host-input">Hostname: </label>
        <input id="host-input" bind:value={host} />
    </div>

    <div class="field">
        <label for="port-input">Port: </label>
        <input id="port-input" bind:value={port} />
    </div>

    <div class="field">
        <label for="user-input">User: </label>
        <input id="user-input" bind:value={user} />
    </div>

    <div class="field">
        <label for="database-input">Database: </label>
        <input id="database-input" bind:value={database} />
    </div>
    <button on:click={handleCreateNewConnection} disabled={loading}>
        Create
    </button>
</div>

<style lang="scss">
    .create-new-connection-form {
        display: grid;

        grid-template-columns: 50% 50%;
        gap: 0.25em;
        width: 50%;
        margin: auto;
    }

    .field {
        display: flex;
        flex-direction: column;
        align-items: start;
    }
</style>
