<script lang="ts">
    import ViewController from "../components/viewController/ViewController.svelte";
    import SidePanel from "../components/sidePanel/SidePanel.svelte";
    import Titlebar from "../components/titlebar/Titlebar.svelte";
    import Toolsbar from "../components/toolbar/Toolbar.svelte";
    import FooterBar from "../components/footerBar/FooterBar.svelte";
</script>

<div class="layout">
    <Titlebar />
    <Toolsbar />
    <SidePanel />
    <ViewController />
    <FooterBar />
</div>

<style lang="scss">
    .layout {
        height: 100vh;

        display: grid;

        grid-template-columns: clamp(20%, 30%, 30em) 1fr;
        grid-template-rows: auto auto 1fr auto;
        height: 100%;
        width: 100%;
    }
</style>
