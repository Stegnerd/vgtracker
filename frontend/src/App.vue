<script setup lang="ts">
import { onMounted } from 'vue';
import { GetConfig } from '../wailsjs/go/controllers/ConfigController';
import { useLayout } from './composables/layout';

  // This starter template is using Vue 3 <script setup> SFCs
  // Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
  //import Greet from "./components/Greet.vue";
  // function toggleDark() {
  //   const element = document.querySelector("html");
  //   element.classList.toggle("app-dark");
  // }
  const {setDefaultValues} = useLayout()

  onMounted(async () => {
    GetConfig().then((result) => {
      // loading in the persisted theme settings for the composable
      setDefaultValues(result.theme.primaryColor, result.theme.surfaceColor, result.theme.preset, result.theme.IsDarkTheme)
    })
  })

</script>

<template>
  <!-- <Button label="Primary" />
  <Button label="Secondary" severity="secondary" />
  <Button label="Success" severity="success" />
  <Button label="Info" severity="info" />
  <Button label="Warn" severity="warn" />
  <Button label="Help" severity="help" />
  <Button label="Danger" severity="danger" />
  <Button label="Contrast" severity="contrast" /> -->
  <RouterView />
  <DynamicDialog />
  <Toast />
  <Toast
    position="bottom-left"
    group="bl"
  />
  <Toast
    position="bottom-right"
    group="br"
  />
</template>

<style scoped></style>
