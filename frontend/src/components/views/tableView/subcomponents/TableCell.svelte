<script lang="ts">
	export let value: any = undefined;
</script>

<td class="table-cell"> {value} </td>

<style lang="scss">
	.table-cell {
		padding: 0.25em 1em;
		white-space: nowrap;
		overflow-x: auto;
		-webkit-user-select: none;

		&:not(:last-child) {
			border-bottom: var(--border-outline-soft);
		}

		&:hover {
			background: var(--hover);
		}
	}
</style>
