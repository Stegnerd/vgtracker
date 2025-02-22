<script setup lang="ts">
  import { useDialog } from "primevue/usedialog";
import { ref } from "vue";
import { useLayout } from "../composables/layout";
import AppThemeConfigurator from "./AppThemeConfigurator.vue";

  const { toggleDarkMode } = useLayout();

  const dialog = useDialog();

  const isDarkMode = ref(true);

  // function toggleDarkMode() {
  //   const element = document.querySelector("html");
  //   element!.classList.toggle("app-dark");
  //   isDarkMode.value = !isDarkMode.value;
  // }

  function openThemeDialog() {
    dialog.open(AppThemeConfigurator, {
      props: {
        modal: false,
        position: "topright"
      }
    });
  }
</script>
<template>
  <div class="topbar-container flex flex-row justify-between items-center pb-10">
    <div />
    <div class="flex justify-center">
      <InputGroup class="min-w-lg">
        <InputText placeholder="Search" />
        <InputGroupAddon>
          <i class="i-mdi-magnify" />
        </InputGroupAddon>
        <!-- <Button icon="i-mdi-magnify" /> -->
      </InputGroup>
    </div>
    <div class="flex gap-6">
      <Button
        :icon="isDarkMode ? 'i-mdi-weather-night' : 'i-mdi-white-balance-sunny'"
        @click="toggleDarkMode()"
      />
      <Button
        icon="i-mdi-palette"
        
        @click="openThemeDialog()"
      />
    </div>
  </div>
</template>
<style></style>
