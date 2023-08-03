<script lang="ts">
	export let name: string = "";
	export let svg: string = "";
	let className: string = "";
	export { className as class }

	$: iconClass = typeof className === "string" ? `icon ${className}` : "icon material-symbols-outlined";

	interface $$Props extends svelteHTML.HTMLAttributes<HTMLSpanElement> {

	}

</script>

<span class={iconClass} {...$$restProps}>
	{#if svg}
		<img src={svg} alt={name} />
	{:else if name}
		{name}
	{:else}
		<slot {...$$restProps} />
	{/if}
</span>

<style lang="scss">
	.icon {
		font-size: inherit;
		line-height: 1;
	}
</style>
