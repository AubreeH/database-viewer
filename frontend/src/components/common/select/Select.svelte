<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import type { Options } from "./types";

	export let options: Options;
	export let props: any = undefined;
	export let value: string = undefined;

	type SelectEvent = Event & { currentTarget: EventTarget & HTMLSelectElement };

	const dispatch = createEventDispatcher();

	function handleSelect(e: SelectEvent) {
		console.log(e);
		if (e?.currentTarget?.value) {
			dispatch("select", getOption(e.currentTarget.value));
		}
	}

	function getOption(value: string) {
		return options.find((o) => o.value === value);
	}
</script>

<select
	class="primary"
	placeholder="Please select a value..."
	{...props}
	on:change={handleSelect}
	bind:value
>
	{#if Array.isArray(options)}
		{#if !value}
			<option value="" disabled selected> Please select a value... </option>
		{/if}
		{#each options as option}
			<option title={option.title} value={option.value}>
				{option.display}
			</option>
		{/each}
	{/if}
</select>
