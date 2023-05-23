<script lang="ts">
    import { Greet } from "../wailsjs/go/main/App.js";
    import theme from "./actions/theme";
    import KeybindManager from "./components/common/keybindManager/KeybindManager.svelte";
    import Modal from "./components/modal/Modal.svelte";
    import Layout from "./layout/Layout.svelte";

    let resultText: string = "Please enter your name below 👇";
    let name: string;
</script>

<svelte:body use:theme />

<svelte:head>
    <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@40,400,0,0"
    />
</svelte:head>

<main>
    <KeybindManager>
        <Layout />
        <Modal />
    </KeybindManager>
</main>

<style>
    main {
        height: 100%;
        width: 100%;
    }
</style>
