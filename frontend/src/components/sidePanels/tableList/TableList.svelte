<script lang="ts">
    import TableListItem from "./TableListItem.svelte";
    import type { IDatabaseTable } from "./types";

    export let tables: IDatabaseTable[] = undefined;
</script>

{#if Array.isArray(tables)}
    {#each tables as table}
        <TableListItem {table} />
    {/each}
{/if}
