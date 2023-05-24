<script lang="ts">
	import {
		Quit,
		WindowMinimise,
		WindowToggleMaximise,
	} from "../../../wailsjs/runtime/runtime";

	import { handleTitlebarDoubleClick } from "./functions";
</script>

<div class="titlebar" on:dblclick={handleTitlebarDoubleClick}>
	<div class="titlebar-start" />
	<div class="titlebar-end">
		<button class="titlebar-button hoverable icon primary" on:click={Quit}>
			close
		</button>
		<button
			class="titlebar-button hoverable icon primary"
			on:click={WindowToggleMaximise}
		>
			maximize
		</button>
		<button
			class="titlebar-button hoverable icon primary"
			on:click={WindowMinimise}
		>
			minimize
		</button>
	</div>
</div>

<style lang="scss">
	.titlebar {
		grid-column: 1/-1;
		grid-row: 1;
		height: 2em;
		border-bottom: 1px solid var(--divider);
		--wails-draggable: drag;

		display: flex;
		justify-content: space-between;

		&,
		* {
			user-select: none;
			-webkit-user-select: none;
		}
	}

	.titlebar-start,
	.titlebar-end {
		display: flex;
		width: fit-content;
		pointer-events: none;
		& > * {
			pointer-events: all;
		}
	}

	.titlebar-end {
		flex-direction: row-reverse;
	}

	.titlebar-button {
		min-height: calc(2em - 1px);
		max-height: calc(2em - 1px);
		padding: 0;
		display: flex;
		align-items: center;
		justify-content: center;
		border: none;
		border-radius: 0;
		box-shadow: none;
		background: none;
	}
</style>
