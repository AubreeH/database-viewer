<script lang="ts">
	import { Greet } from "../wailsjs/go/main/App.js";
	import theme from "./actions/theme";
	import KeybindManager from "./components/common/keybindManager/KeybindManager.svelte";
	import Modal from "./components/modal/Modal.svelte";
	import NotificationsContext from "./components/notifications/NotificationsContext.svelte";
	import Layout from "./layout/Layout.svelte";

	let resultText: string = "Please enter your name below 👇";
	let name: string;
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
