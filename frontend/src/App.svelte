<script lang="ts">
	import theme from "./actions/theme";
	import KeybindManager from "./components/common/keybindManager/KeybindManager.svelte";
	import Modal from "./components/modal/Modal.svelte";
	import NotificationsContext from "./components/notifications/NotificationsContext.svelte";
	import Layout from "./layout/Layout.svelte";
</script>

<svelte:body use:theme />

<main>
	<NotificationsContext>
		<KeybindManager>
			<Layout />
			<Modal />
		</KeybindManager>
	</NotificationsContext>
</main>

<style>
	main {
		height: 100%;
		width: 100%;
	}
</style>
