<script lang="ts">
    import ThemeSelect from "$lib/theme/theme-select.svelte";
    import { Search, Library, Star, ChartLine } from "@lucide/svelte";
    import { page } from "$app/state";
</script>

<div class="navbar-start">
    <div class="dropdown">
        <div tabindex="0" role="button" class="btn btn-ghost lg:hidden">
            <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M4 6h16M4 12h8m-8 6h16"
                />
            </svg>
        </div>
        <ul
            tabindex="-1"
            class="menu menu-sm dropdown-content bg-base-200 rounded-box z-1 mt-3 w-52 p-2 shadow"
        >
            <li>
                <a
                    href="/search"
                    class={page.url.pathname === "/search" ? "menu-active" : ""}
                >
                    <Search class=" h-4" />Search</a
                >
            </li>
            <li>
                <a
                    href="/"
                    class={page.url.pathname === "/" ? "menu-active" : ""}
                >
                    <ChartLine class=" h-4" />Stats</a
                >
            </li>
            <li>
                <a
                    href="/library"
                    class={page.url.pathname === "/library"
                        ? "menu-active"
                        : ""}
                >
                    <Library class=" h-4" />Library</a
                >
            </li>
            <li>
                <a
                    href="/wishlist"
                    class={page.url.pathname === "/wishlist"
                        ? "menu-active"
                        : ""}
                >
                    <Star class=" h-4" />Wishlist</a
                >
            </li>
        </ul>
    </div>
    <span class="primary text-xl">VGTracker</span>
</div>
<div class="navbar-center hidden lg:flex">
    <label class="input">
        <Search />
        <input type="search" class="grow" placeholder="Search" />
        <kbd class="kbd kbd-sm">⌘</kbd>
        <kbd class="kbd kbd-sm">K</kbd>
    </label>
</div>
<div class="navbar-end">
    <ThemeSelect></ThemeSelect>
</div>
