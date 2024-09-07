<script setup lang="ts">
  import { ref } from "vue";

  const isDarkMode = ref(true);

  function toggleDarkMode() {
    const element = document.querySelector("html");
    element!.classList.toggle("app-dark");
    isDarkMode.value = !isDarkMode.value;
  }
</script>
<template>
  <div class="topbar-container flex flex-row justify-between items-center">
    <div class="tester"></div>
    <div class="flex justify-center">
      <InputGroup class="min-w-lg">
        <InputText placeholder="Search" />
        <Button icon="i-mdi-magnify" />
      </InputGroup>
    </div>
    <div class="flex gap-6">
      <Button
        :icon="isDarkMode ? 'i-mdi-weather-night' : 'i-mdi-white-balance-sunny'"
        @click="toggleDarkMode()"
      />
      <Button icon="i-mdi-palette" class="m-r" />
    </div>
  </div>
</template>
<style></style>
