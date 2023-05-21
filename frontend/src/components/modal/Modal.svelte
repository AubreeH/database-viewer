<script lang="ts">
    import { fade } from "svelte/transition";
    import ModalStore, { closeModal } from "../../stores/modal";

    function handleClickModal(e: MouseEvent | KeyboardEvent) {
        e.stopPropagation();
    }

    function handleClickModalBackground(e: MouseEvent | KeyboardEvent) {
        closeModal();
    }
</script>

{#if $ModalStore?.open}
    <div
        transition:fade={{ duration: 100 }}
        class="modal-background"
        on:click={handleClickModalBackground}
        on:keypress={handleClickModalBackground}
    >
        <div
            class="modal"
            on:click={handleClickModal}
            on:keypress={handleClickModal}
        >
            <svelte:component
                this={$ModalStore.component}
                {...$ModalStore.props ?? {}}
            />
        </div>
    </div>
{/if}

<style lang="scss">
    .modal-background {
        position: fixed;

        top: 0;
        bottom: 0;
        left: 0;
        right: 0;

        display: flex;
        justify-content: center;
        align-items: center;
        background: #0006;
    }

    .modal {
        min-height: 40%;
        min-width: 40%;

        background: var(--primary);
    }
</style>
