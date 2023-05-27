<script lang="ts">
	import { onMount } from "svelte";
	import { GetPasswordBehaviours, SaveConnection } from "../../../../wailsjs/go/connectionsBinding/ConnectionsBinding";
	import { connections } from "../../../../wailsjs/go/models";
	import { WindowReload } from "../../../../wailsjs/runtime/runtime";
	import { closeModal } from "../../../stores/modal";
	import Select from "../../common/select/Select.svelte";
	import { PasswordBehaviour } from "./types";
	import type { IOption } from "../../common/select/types";
	import { DatabaseDriverStringMap, EDatabaseDrivers } from "../../../types/databaseDrivers";
	import { get } from "svelte/store";

	export let connection: connections.Connection = undefined;

	let name: string = "";
	let host: string = "";
	let port: string = "";
	let user: string = "";
	let database: string = "";
	let password: string = "";
	let selectedBehaviour: PasswordBehaviour = PasswordBehaviour.EmptyString;
	let selectedDriver: string = "";

	let behaviours: connections.PasswordBehaviourDescription[] = undefined;
	let drivers: IOption[] = undefined;

	$: setVariables(connection);

	function setVariables(connection: connections.Connection) {
		if (connection) {
			name = connection.name;
			host = connection.host;
			port = connection.port;
			user = connection.user;
			database = connection.database;
			selectedBehaviour = connection.password_save_type as PasswordBehaviour;
			selectedDriver = connection.driver_name;
			password = connection.password;
		}
	}

	onMount(async () => {
		behaviours = await GetPasswordBehaviours();
		drivers = Array.from(DatabaseDriverStringMap.values());
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
				driver: (selectedDriver ?? "").split("???", 2)[0],
				driver_name: selectedDriver ?? "",
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
</script>

<form class="create-new-connection-form" on:submit={handleSaveConnection}>
	<fieldset class="field">
		<label for="name-input">Connection Name: </label>
		<input id="name-input" class="primary" bind:value={name} />
	</fieldset>

	<fieldset class="field">
		<label for="database-driver-input">Driver: </label>
		<Select options={drivers} props={{ id: "database-driver-input" }} bind:value={selectedDriver} />
	</fieldset>

	{#if selectedDriver === DatabaseDriverStringMap.get(EDatabaseDrivers.SQLite).value}
		<fieldset class="field">
			<label for="host-input">File Path: </label>
			<input id="host-input" class="primary" bind:value={host} />
		</fieldset>
	{:else if selectedDriver}
		<fieldset class="field">
			<label for="host-input">Hostname: </label>
			<input id="host-input" class="primary" bind:value={host} />
		</fieldset>

		<fieldset class="field">
			<label for="port-input">Port: </label>
			<input id="port-input" class="primary" bind:value={port} />
		</fieldset>

		<fieldset class="field">
			<label for="user-input">User: </label>
			<input id="user-input" class="primary" bind:value={user} />
		</fieldset>

		<fieldset class="field">
			<label for="database-input">Database: </label>
			<input id="database-input" class="primary" bind:value={database} />
		</fieldset>

		<fieldset class="field">
			<label for="password-save-type-input">Password Save Type: </label>
			<Select options={getOptions(behaviours)} props={{ id: "password-save-type-input" }} bind:value={selectedBehaviour} />
		</fieldset>

		{#if selectedBehaviour === PasswordBehaviour.SaveAsPlainText}
			<fieldset class="field">
				<label for="password-input">Password: </label>
				<input id="password-input" class="primary" bind:value={password} type="password" />
			</fieldset>
		{/if}
	{/if}

	<button type="submit" class="primary" disabled={loading}> Save </button>
</form>

<style lang="scss">
	.create-new-connection-form {
		display: flex;
		flex-direction: column;
		gap: 1em;
		margin: auto;
	}

	.field {
		width: 25rem;
		display: flex;
		flex-direction: column;
		align-items: start;
		border: none;
		padding: 0;
	}
</style>
