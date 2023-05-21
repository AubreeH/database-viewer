<script lang="ts">
    export let name: string = "";
    export let props: any = undefined;

    $: iconClass = typeof props?.class === "string" ? `icon ${props.class}` : "icon material-symbols-outlined";
</script>

<span class={iconClass} {...props}>
    {#if name}
        {name}
    {:else}
        <slot />
    {/if}
</span>

<style lang="scss">
    .icon {
        font-size: inherit;
        line-height: 1em;
    }
</style>