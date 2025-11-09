<script lang="ts">
    import { onMount } from 'svelte';
    import { themes } from './themes';

    onMount(() => {
        const localStorageValue = localStorage.getItem('theme') ?? 'retro';

        const correctInputField = document.getElementById(
            `${localStorageValue}-theme-id`,
        );
        if (correctInputField) {
            (correctInputField as HTMLInputElement).checked = true;
        }
    });

    function select(value: Event) {
        if (value) {
            localStorage.setItem(
                'theme',
                (value.target as HTMLInputElement).value,
            );
        }
    }
</script>

<div class="dropdown">
    <div tabindex="0" role="button" class="btn m-1">
        Theme
        <svg
            width="12px"
            height="12px"
            class="inline-block h-2 w-2 fill-current opacity-60"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 2048 2048"
        >
            <path d="M1799 349l242 241-1017 1017L7 590l242-241 775 775 775-775z"
            ></path>
        </svg>
    </div>
    <ul
        tabindex="-1"
        class="dropdown-content bg-base-300 rounded-box z-1 w-52 p-2 shadow-2xl"
    >
        <li>
            <input
                type="radio"
                name="theme-dropdown"
                class="theme-controller w-full btn btn-sm btn-block btn-ghost justify-start"
                aria-label="default"
                value="default"
            />
        </li>
        {#each themes as theme (theme)}
            <li>
                <input
                    id="{theme}-theme-id"
                    type="radio"
                    name="theme-dropdown"
                    class="theme-controller w-full btn btn-sm btn-block btn-ghost justify-start"
                    aria-label={theme}
                    value={theme}
                    oninput={(v) => select(v)}
                />
            </li>
        {/each}
    </ul>
</div>
