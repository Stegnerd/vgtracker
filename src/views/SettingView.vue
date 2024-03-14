<script setup lang="ts">
import {useConfigStore} from "@/stores/configStore.ts";

const store = useConfigStore();
const s = store.configuration;
console.warn('config from store:', s)

</script>
<template>
  <p>setting</p>
</template>
