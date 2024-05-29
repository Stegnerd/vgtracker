<script setup lang="ts">
import LayoutDefault from "@/layouts/LayoutDefault.vue";

import { useConfigStore } from "./stores/configStore.ts";

const configStore = useConfigStore();

configStore.getConfig().then(() => {
  const conf = configStore.configuration;
  const root = document.getElementsByTagName("html")[0];
  if (conf?.theme === "Dark") {
    root.classList.toggle("dark");
  }
});
</script>

<template>
  <LayoutDefault />
</template>

<style scoped></style>
