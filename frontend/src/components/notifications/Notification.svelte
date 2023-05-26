<script lang="ts">
	import { fly, slide } from "svelte/transition";
	import Icon from "../common/icon/Icon.svelte";
	import type { INotification } from "./types";

	export let notification: INotification;
	export let id: string;
</script>

<div
	class="notification"
	{id}
	role="alertdialog"
	style="--bg: {notification.background}; --fg: {notification.color}"
	in:fly={{ x: 200 }}
	out:slide
	on:click
	on:keypress
>
	{#if notification.icon}
		<Icon {...notification.iconProps}>
			{notification.icon}
		</Icon>
	{/if}
	<span>
		{notification.message}
	</span>
</div>

<style lang="scss">
	.notification {
		display: grid;
		grid-template-columns: auto 1fr;
		position: relative;
		max-width: 25em;
		background-color: var(--bg, white);
		color: var(--fg, black);
		align-items: center;
		gap: 1em;
		padding: 0.5em 1em;
		border-radius: 0.5em;
		margin-bottom: 1em;
		overflow: hidden;
		font-weight: 500;
		min-height: fit-content;

		:global(.icon) {
			font-size: 1.5em;
			font-weight: 500;
		}
	}
</style>
