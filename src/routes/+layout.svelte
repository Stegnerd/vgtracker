<script lang="ts">
    import "../app.css";
    import ThemeSelect from "$lib/theme/theme-select.svelte";
    import { Search, Library, Star } from "@lucide/svelte";
    let { children } = $props();
</script>

<div class="app h-dvh flex flex-col">
    <div class="navbar bg-base-100 shadow-sm">
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
                    <li><a>Item 1</a></li>
                    <li>
                        <a>Parent</a>
                        <ul class="p-2">
                            <li><a>Submenu 1</a></li>
                            <li><a>Submenu 2</a></li>
                        </ul>
                    </li>
                    <li><a>Item 3</a></li>
                </ul>
            </div>
            <span class="primary text-xl">VGTracker</span>
        </div>
        <div class="navbar-center hidden lg:flex">
            <!-- <ul class="menu menu-horizontal px-1">
                <li><a>Item 1</a></li>
                <li>
                    <details>
                        <summary>Parent</summary>
                        <ul class="p-2">
                            <li><a>Submenu 1</a></li>
                            <li><a>Submenu 2</a></li>
                        </ul>
                    </details>
                </li>
                <li><a>Item 3</a></li>
            </ul> -->
            <label class="input">
                <!-- <svg
                    class="h-[1em] opacity-50"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                >
                    <g
                        stroke-linejoin="round"
                        stroke-linecap="round"
                        stroke-width="2.5"
                        fill="none"
                        stroke="currentColor"
                    >
                        <circle cx="11" cy="11" r="8"></circle>
                        <path d="m21 21-4.3-4.3"></path>
                    </g>
                </svg> -->

                <Search />
                <input type="search" class="grow" placeholder="Search" />
                <kbd class="kbd kbd-sm">⌘</kbd>
                <kbd class="kbd kbd-sm">K</kbd>
            </label>
        </div>
        <div class="navbar-end">
            <ThemeSelect></ThemeSelect>
        </div>
    </div>
    <div class="flex flex-col grow">
        <div class="flex flex-row m-4 h-full">
            <div class="hidden lg:flex lg:w-48 lg:flex-none mr-4 size-fit">
                <!-- <div class="rounded-box bg-primary text-primary-content w-64 p-4">
                <div class="flex flex-col">
                    <p>item 1</p>
                    <p>item 2</p>
                </div>
            </div> -->

                <ul class="menu bg-base-200 w-48">
                    <li>
                        <a> <Search class=" h-4" />Search</a>
                    </li>
                    <li>
                        <a> <Library class=" h-4" />Library</a>
                    </li>
                    <li>
                        <a> <Star class=" h-4" />Wishlist</a>
                    </li>
                    <li><a class="menu-active">Item 2</a></li>
                </ul>
            </div>
            <div class="flex-auto flex-col bg-base-200 p-[0.5rem]">
                {@render children()}
            </div>
        </div>
    </div>
</div>

<style>
</style>
