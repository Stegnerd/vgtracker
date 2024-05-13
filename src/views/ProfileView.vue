<script setup lang="ts">
  import { useConfigStore } from "@/stores/configStore.ts";

  import { UpdateConfigInput } from "../../src-tauri/bindings/config/UpdateConfigInput.ts";

  const store = useConfigStore();
  async function clicker() {
    const input = {
      twitchClientId: "cat",
      twitchClientSecret: "dog"
    } as UpdateConfigInput;
    await store.updateConfig(input);
  }
</script>

<template>
  Profile
  <PrimeButton
    label="Submit"
    type="submit"
    data-testid="submit-button"
    @click="clicker()"
  />
</template>
