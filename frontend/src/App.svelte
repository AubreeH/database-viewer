<script lang="ts">
  import {Greet} from '../wailsjs/go/main/App.js'
    import Layout from './layout/Layout.svelte';

  let resultText: string = "Please enter your name below 👇"
  let name: string

  function greet(): void {
    Greet(name).then(result => resultText = result)
  }
</script>

<head>
  <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@40,400,0,0" />
</head>
<main>
  <Layout />
</main>

<style>
  main {
    height: 100%;
    width: 100%;
  }

</style>
