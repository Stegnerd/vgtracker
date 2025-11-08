<script lang="ts">
    import { Search, Library, Star, ChartLine } from "@lucide/svelte";
    import { page } from "$app/state";
</script>

<div class="hidden lg:flex lg:w-48 lg:flex-none mr-4 size-fit">
    <ul class="menu bg-base-200 w-48">
        <li>
            <a
                href="/search"
                class={page.url.pathname === "/search" ? "menu-active" : ""}
            >
                <Search class=" h-4" />Search</a
            >
        </li>
        <li>
            <a href="/" class={page.url.pathname === "/" ? "menu-active" : ""}>
                <ChartLine class=" h-4" />Stats</a
            >
        </li>
        <li>
            <a
                href="/library"
                class={page.url.pathname === "/library" ? "menu-active" : ""}
            >
                <Library class=" h-4" />Library</a
            >
        </li>
        <li>
            <a
                href="/wishlist"
                class={page.url.pathname === "/wishlist" ? "menu-active" : ""}
            >
                <Star class=" h-4" />Wishlist</a
            >
        </li>
    </ul>
</div>
