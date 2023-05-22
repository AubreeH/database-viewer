<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import { EventsEmit, EventsOn } from "../../../../wailsjs/runtime/runtime";
    import { fade } from "svelte/transition";

    export let id: string;

    export let options: { display: string; handler?: () => void }[] = [];

    let x: number = undefined;
    let y: number = undefined;
    let show: boolean = false;
    let element: HTMLElement = undefined;

    $: dataTag = `${id}-context-menu`;

    $: contextmenu = {
        "data-context-menu": dataTag,
    };

    function onContextMenu(e: MouseEvent) {
        if (e.ctrlKey) {
            return;
        }

        const element = e?.target as HTMLElement;
        if (element && element?.attributes) {
            const attributes = element?.attributes;
            const tag = attributes?.getNamedItem("data-context-menu")?.value;

            if (tag && tag === dataTag) {
                handleContextMenu(e);
            }
        }
    }

    function handleContextMenu(e: MouseEvent) {
        e.preventDefault();
        EventsEmit("close-context-menu");
        x = e.clientX;
        y = e.clientY;
        show = true;

        document.addEventListener("click", handleClickAway);
    }

    function handleClickAway(e: MouseEvent) {
        if (e.target) {
            const target = e.target as HTMLElement;
            if (!target.closest(`#${dataTag}`)) {
                show = false;
                document.removeEventListener("click", handleClickAway);
            }
        }
    }

    function handleSelect(handler?: () => void) {
        show = false;
        if (handler) {
            handler();
        }
    }

    onMount(() => {
        EventsOn("close-context-menu", () => (show = false));
        document.addEventListener("contextmenu", onContextMenu);
    });

    onDestroy(() => {
        document.removeEventListener("contextmenu", onContextMenu);
    });
</script>

<slot {contextmenu} />

{#if show}
    <div
        transition:fade={{ duration: 50 }}
        bind:this={element}
        id={dataTag}
        class={"context-menu"}
        style={`top: ${y}px; left: ${x}px`}
    >
        {#each options as option}
            <button
                class="primary hoverable"
                on:click={() => handleSelect(option.handler)}
            >
                {option.display}
            </button>
        {/each}
    </div>
{/if}

<style lang="scss">
    .context-menu {
        z-index: 100;
        background: var(--primary);
        min-width: 100px;
        height: fit-content;
        position: fixed;
        display: flex;
        flex-direction: column;
        border-radius: 0.5em;
        border: var(--border-outline);
        box-shadow: var(--light-box-shadow);

        button {
            margin: 0;
            border-radius: 0;
            border: none;
            box-shadow: none;

            &:first-child {
                border-top-left-radius: 0.5em;
                border-top-right-radius: 0.5em;
            }

            &:last-child {
                border-bottom-left-radius: 0.5em;
                border-bottom-right-radius: 0.5em;
            }
        }
    }
</style>
