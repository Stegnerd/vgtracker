<script lang="ts">
    import '../app.css';
    import NavBar from '../layout/nav-bar.svelte';
    import SideNav from '../layout/side-nav.svelte';
    import EnvironmentConfigModal from '$lib/environment-config-modal.svelte';

    let { data, children } = $props();
</script>

<div class="app h-dvh flex flex-col">
    <div class="navbar bg-base-100 shadow-sm">
        <NavBar></NavBar>
    </div>
    <div class="flex flex-col grow">
        <div class="flex flex-row m-4 h-full">
            <SideNav></SideNav>
            <div class="flex-auto flex-col bg-base-200 p-2">
                <EnvironmentConfigModal isTokenSet={data.isTokenSet}
                ></EnvironmentConfigModal>
                {@render children()}
            </div>
        </div>
    </div>
</div>

<style>
</style>
