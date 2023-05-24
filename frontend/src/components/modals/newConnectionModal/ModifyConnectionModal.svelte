<script lang="ts">
	import { onMount } from "svelte";
	import {
		GetPasswordBehaviours,
		SaveConnection,
	} from "../../../../wailsjs/go/connectionsBinding/ConnectionsBinding";
	import { connections } from "../../../../wailsjs/go/models";
	import { WindowReload } from "../../../../wailsjs/runtime/runtime";
	import { closeModal } from "../../../stores/modal";
	import Select from "../../common/select/Select.svelte";
	import { PasswordBehaviour } from "./types";
	import type { IOption } from "src/components/common/select/types";

	export let connection: connections.Connection = undefined;

	let name: string = "";
	let host: string = "";
	let port: string = "";
	let user: string = "";
	let database: string = "";
	let password: string = "";
	let selectedBehaviour: PasswordBehaviour = PasswordBehaviour.EmptyString;

	let behaviours: connections.PasswordBehaviourDescription[] = undefined;

	$: setVariables(connection);

	function setVariables(connection: connections.Connection) {
		if (connection) {
			name = connection.name;
			host = connection.host;
			port = connection.port;
			user = connection.user;
			database = connection.database;
			selectedBehaviour = connection.password_save_type as PasswordBehaviour;
			password = connection.password;
		}
	}

	onMount(async () => {
		behaviours = await GetPasswordBehaviours();
	});

	let loading = false;

	async function handleSaveConnection() {
		loading = true;

		await SaveConnection(
			new connections.Connection({
				name,
				host,
				port,
				user,
				database,
				password_save_type: selectedBehaviour,
				password,
			})
		);

		WindowReload();
		closeModal();
	}

	function getOptions(behaviours: connections.PasswordBehaviourDescription[]) {
		if (Array.isArray(behaviours)) {
			return behaviours.map((b) => ({
				display: b.display,
				value: b.behaviour,
				title: b.description,
			}));
		}

		return [];
	}

	function handleSelect(e: CustomEvent<IOption>) {
		console.log(e);
		selectedBehaviour = e.detail.value as PasswordBehaviour;
	}

	$: console.log(selectedBehaviour, connection);
</script>

<div class="create-new-connection-form">
	<div class="field">
		<label for="name-input">Connection Name: </label>
		<input id="name-input" class="primary" bind:value={name} />
	</div>

	<div class="field">
		<label for="host-input">Hostname: </label>
		<input id="host-input" class="primary" bind:value={host} />
	</div>

	<div class="field">
		<label for="port-input">Port: </label>
		<input id="port-input" class="primary" bind:value={port} />
	</div>

	<div class="field">
		<label for="user-input">User: </label>
		<input id="user-input" class="primary" bind:value={user} />
	</div>

	<div class="field">
		<label for="database-input">Database: </label>
		<input id="database-input" class="primary" bind:value={database} />
	</div>

	<div class="field">
		<label for="password-save-type-input">Password Save Type: </label>
		<Select
			options={getOptions(behaviours)}
			props={{ id: "password-save-type-input" }}
			bind:value={selectedBehaviour}
		/>
	</div>

	{#if selectedBehaviour === PasswordBehaviour.SaveAsPlainText}
		<div class="field">
			<label for="password-input">Password: </label>
			<input
				id="password-input"
				class="primary"
				bind:value={password}
				type="password"
			/>
		</div>
	{/if}

	<button on:click={handleSaveConnection} class="primary" disabled={loading}>
		Save
	</button>
</div>

<style lang="scss">
	.create-new-connection-form {
		display: flex;
		flex-direction: column;
		gap: 0.5em;
		margin: auto;
	}

	.field {
		width: 25rem;
		display: flex;
		flex-direction: column;
		align-items: start;
	}
</style>
