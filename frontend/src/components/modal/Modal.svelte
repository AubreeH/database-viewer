<script lang="ts">
	import { fade } from "svelte/transition";
	import ModalStore, { ModalSize, closeModal } from "../../stores/modal";

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
		class={`modal-background ${$ModalStore.size ?? ModalSize.Medium}`}
		on:click={handleClickModalBackground}
		on:keypress={handleClickModalBackground}
	>
		<div
			class="modal"
			on:click={handleClickModal}
			on:keypress={handleClickModal}
		>
			{#if $ModalStore.title}
				<h4 class="modal-title">{$ModalStore.title}</h4>
			{/if}

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
		width: fit-content;
		height: fit-content;
		padding: 1.5em 2em 2em;
		border-radius: 1em;

		background: var(--primary);
	}

	.modal-title {
		width: 100%;
		padding: 0 0 0.5em 0;
		margin: 0 0 1em 0;
		border-bottom: 1px solid var(--outline);
	}
</style>
