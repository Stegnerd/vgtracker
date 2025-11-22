<script lang="ts">
  import { invoke } from '@tauri-apps/api/core';
  import { onMount } from 'svelte';

  let { isTokenSet } = $props();

  onMount(() => {
    console.log('the component has mounted');
    if (!isTokenSet) {
      const tokenModal = document.getElementById(
        'token_modal',
      ) as HTMLDialogElement;
      console.warn('tokemModal', tokenModal);
      tokenModal?.showModal();
    }
  });

  function refreshToken() {
    invoke('refresh_token').then((result) => {
      if (result) {
        window.location.reload();
      }
    });
  }
</script>

<dialog id="token_modal" class="modal">
  <div class="modal-box">
    <h3 class="text-lg font-bold">Welcome to VGTracker!</h3>
    <p class="py-4">
      It looks like your environment variables are not set correctly
    </p>

    <ul class="py-4 px-4">
      <li class="py-2">
        Follow the instructions <a
          class="link link-primary"
          href="https://api-docs.igdb.com/#getting-started"
          target="_blank">here</a
        > to get your Twitch credentials
      </li>
      <li class="py-2">
        then set these variables in your shell environment
        <ul>
          <li><code>vgt_twitch_client_id=your_client_id</code></li>
          <li><code>vgt_twitch_client_secret=your_client_secret</code></li>
        </ul>
      </li>
      <li class="py-2">
        After setting the environment variables, click the button below to
        refresh your token or restart the application
      </li>
    </ul>

    <div class="modal-action">
      <form method="dialog">
        <!-- if there is a button in form, it will close the modal -->
        <button onclick={refreshToken} class="btn">Refresh Token</button>
      </form>
    </div>
  </div>
</dialog>

<style scoped>
  li {
    list-style-type: circle;
  }

  li li {
    list-style-type: unset;
  }
</style>
