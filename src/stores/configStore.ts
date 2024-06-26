import { invoke } from "@tauri-apps/api/tauri";
import { defineStore } from "pinia";
import { ref } from "vue";

import { ReadConfigOutput } from "../../src-tauri/bindings/config/ReadConfigOutput.ts";
import { Theme } from "../../src-tauri/bindings/config/Theme.ts";
import { UpdateConfigInput } from "../../src-tauri/bindings/config/UpdateConfigInput.ts";
import { toastSuccess } from "../composables/useToast.ts";

export const useConfigStore = defineStore("config", () => {
  const configuration = ref<ReadConfigOutput | undefined>();

  async function getConfig(): Promise<void> {
    return await invoke<ReadConfigOutput>("get_user_configuration").then(
      (cfg: ReadConfigOutput) => {
        configuration.value = cfg;
      }
    );
  }

  async function updateConfig(input: UpdateConfigInput) {
    return await invoke<ReadConfigOutput>("update_user_configuration", {
      input
    }).then((cfg: ReadConfigOutput) => {
      configuration.value = cfg;
      toastSuccess("Configration Updated", "");
    });
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
