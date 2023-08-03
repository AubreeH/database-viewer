<script lang="ts">
	import Icon from "../icon/Icon.svelte";
    
    export let offset: number;
    export let lowerBound:number = 0;
    export let total:number;

    function handlePaginate(direction: 1 | -1) {
        let newOffset = offset + direction
        if (lowerBound !== undefined && newOffset < lowerBound) {
            offset = lowerBound
        } else if (total !== undefined && newOffset > (total - 1)) {
            offset = total - 1
        } else {
            offset = newOffset
        }
    }

    $: buttonLeftDisabled = (!!(lowerBound !== undefined && offset <= lowerBound))
    $: buttonRightDisabled = (!!(total !== undefined && offset >= (total - 1)))

    
</script>

<div class="pagination-controls">
    <button 
        class="primary" 
        on:click={() => handlePaginate(-1)} 
        disabled={buttonLeftDisabled}
    >
        <Icon>
            chevron_left
        </Icon>
    </button>
    <span class="pagination-controls-display">
        {offset + 1}
    </span>
    <span class="pagination-controls-display">
        /
    </span>
    <span class="pagination-controls-display">
        {total}
    </span>
    <button 
        class="primary" 
        on:click={() => handlePaginate(1)} 
        disabled={buttonRightDisabled}
    >
        <Icon>
            chevron_right
        </Icon>
    </button>
</div>

<style lang="scss">
    .pagination-controls {
		display: grid;
		grid-template-columns: repeat(5, auto);
		gap: 1em;
		justify-content: center;
		align-content: center;
	}

	.pagination-controls-display {
		display: grid;
		justify-content: center;
		align-content: center;
	}
</style>