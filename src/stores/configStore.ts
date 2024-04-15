import { invoke } from "@tauri-apps/api/tauri";
import { defineStore } from "pinia";
import { ref } from "vue";

import { ReadConfigOutput } from "../../src-tauri/bindings/config/ReadConfigOutput.ts";
import { Theme } from "../../src-tauri/bindings/config/Theme.ts";
import { UpdateConfigInput } from "../../src-tauri/bindings/config/UpdateConfigInput.ts";

export const useConfigStore = defineStore("config", () => {
  const configuration = ref<ReadConfigOutput | undefined>();

  async function getConfig(): Promise<void> {
    return await invoke<ReadConfigOutput>("get_user_config").then(
      (cfg: ReadConfigOutput) => {
        configuration.value = cfg;
      }
    );
  }

  async function updateConfig(input: UpdateConfigInput) {
    return await invoke<ReadConfigOutput>("update_user_config", { input }).then(
      (cfg: ReadConfigOutput) => {
        configuration.value = cfg;
      }
    );
  }

  async function updateTheme(theme: Theme) {
    return await invoke("update_theme", { themeChange: theme }).then(() => {
      if (configuration.value !== undefined) {
        configuration.value.theme = theme;
      }
    });
  }

  return {
    configuration,
    getConfig,
    updateConfig,
    updateTheme
  };
});
