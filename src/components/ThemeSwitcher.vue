<script setup lang="ts">
  import { storeToRefs } from "pinia";
  import { ref } from "vue";

  import { useConfigStore } from "@/stores/configStore.ts";

  const configStore = useConfigStore();
  const { configuration } = storeToRefs(configStore);
  let iconClass = ref(
    configuration.value?.theme === "Light" ? "pi-sun" : "pi-moon"
  );

  async function onThemeToggle() {
    const root = document.getElementsByTagName("html")[0];
    root.classList.toggle("dark");
    const input = configuration.value?.theme === "Light" ? "Dark" : "Light";
    await configStore.updateTheme(input).then(() => {
      iconClass.value =
        configuration.value?.theme === "Light" ? "pi-sun" : "pi-moon";
    });
  }
</script>

<template>
  <div class="flex justify-center">
    <ul class="flex list-none m-0 p-0 gap-2 items-center">
      <li>
        <button
          type="button"
          class="inline-flex border w-48 h-12 p-0 items-center justify-center surface-0 dark:surface-800 border border-surface-200 dark:border-surface-600 rounded"
          @click="onThemeToggle"
        >
          <i :class="`dark:text-white pi ${iconClass}`" />
        </button>
      </li>
    </ul>
  </div>
</template>

<style scoped></style>
