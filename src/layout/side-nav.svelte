<script lang="ts">
    import { Search, Library, Star, ChartLine } from '@lucide/svelte';
    import { page } from '$app/state';
    import { resolve } from '$app/paths';
    // TODO: when we get to the actual game looking up,
    // we need to make the active route be library even though the url name is different
    // same with nav-bar
</script>

<div class="hidden lg:flex lg:w-48 lg:flex-none mr-4 size-fit">
    <ul class="menu bg-base-200 w-48">
        <li>
            <a
                href={resolve('/search')}
                class={page.url.pathname === '/search' ? 'menu-active' : ''}
            >
                <Search class=" h-4" />Search</a
            >
        </li>
        <li>
            <a
                href={resolve('/')}
                class={page.url.pathname === '/' ? 'menu-active' : ''}
            >
                <ChartLine class=" h-4" />Stats</a
            >
        </li>
        <li>
            <a
                href={resolve('/library')}
                class={page.url.pathname === '/library' ? 'menu-active' : ''}
            >
                <Library class=" h-4" />Library</a
            >
        </li>
        <li>
            <a
                href={resolve('/wishlist')}
                class={page.url.pathname === '/wishlist' ? 'menu-active' : ''}
            >
                <Star class=" h-4" />Wishlist</a
            >
        </li>
    </ul>
</div>
