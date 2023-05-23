<script lang="ts">
    import { openTabFocused } from "../../../stores/mainView";
    import TableView from "../../views/tableView/TableView.svelte";
    import type { IDatabaseTable } from "./types";

    export let table: IDatabaseTable;

    function handleClick() {
        openTabFocused({
            component: TableView,
            props: { table },
            name: `Table: ${table.name}`,
        });
    }
</script>

<button class="primary hoverable" on:click={handleClick}>
    {table.name}
</button>
