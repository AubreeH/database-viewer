<script lang="ts">
	import { onMount, onDestroy } from "svelte";
	import { EShadowWeight } from "./types";

	export let weight = EShadowWeight.Light;
	export let size: number = 10;

	let element: HTMLDivElement = undefined;
	let shadow = "";

	onMount(() => {
		element.addEventListener("scroll", () => handleScroll(element));
	});

	onDestroy(() => {
		element.removeEventListener("scroll", () => handleScroll(element));
	});

	function handleScroll(element: HTMLDivElement) {
		if (element) {
			let boxShadow = "";
			let shadowPixelSize = `${size}px`;

			// Left Shadow
			if (element.scrollLeft !== 0) {
				boxShadow += `inset ${shadowPixelSize} 0 ${shadowPixelSize} -${shadowPixelSize} ${weight}`;
			} else {
				boxShadow += `inset 0 0 0 0 transparent`;
			}

			// Right Shadow
			if (element.scrollLeft + element.clientWidth !== element.scrollWidth) {
				boxShadow += `, inset -${shadowPixelSize} 0 ${shadowPixelSize} -${shadowPixelSize} ${weight}`;
			} else {
				boxShadow += `, inset 0 0 0 0 transparent`;
			}

			// Top Shadow
			if (element.scrollTop !== 0) {
				boxShadow += `, inset 0 ${shadowPixelSize} ${shadowPixelSize} -${shadowPixelSize} ${weight}`;
			} else {
				boxShadow += `, inset 0 0 0 0 transparent`;
			}

			// Bottom Shadow
			if (element.scrollTop + element.clientHeight !== element.scrollHeight) {
				boxShadow += `, inset 0 -${shadowPixelSize} ${shadowPixelSize} -${shadowPixelSize} ${weight}`;
			} else {
				boxShadow += `, inset 0 0 0 0 transparent`;
			}

			shadow = boxShadow;
		}
	}
</script>

<div class="hidden-scrollbar-container" style="--hidden-scrollbar-shadow: {shadow}">
	<div class="hidden-scrollbar-container-inner hidden-scrollbar" bind:this={element}>
		<slot />
	</div>
</div>

<style lang="scss">
	.hidden-scrollbar-container {
		position: relative;
		height: 100%;
		width: 100%;
		overflow: auto;

		&::after {
			transition: box-shadow var(--speed4);
			content: "";
			position: absolute;
			top: 0;
			bottom: 0;
			left: 0;
			right: 0;
			pointer-events: none;
			box-shadow: var(--hidden-scrollbar-shadow);
		}
	}

	.hidden-scrollbar-container-inner {
		height: 100%;
		width: 100%;
		overflow: auto;
	}
</style>
