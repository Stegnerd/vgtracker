import {defineStore} from "pinia";
import {invoke} from "@tauri-apps/api/tauri";
import {ref} from "vue";
import {UpdateConfigInput} from "../../src-tauri/bindings/config/UpdateConfigInput.ts";
import {ReadConfigOutput} from "../../src-tauri/bindings/config/ReadConfigOutput.ts";

export const useConfigStore = defineStore('config', () => {
    const configuration = ref<ReadConfigOutput | undefined>();

    async function getConfig(): Promise<void> {
        return await invoke<ReadConfigOutput>('get_user_config').then((cfg: ReadConfigOutput) => {
            console.warn('in store', cfg)
            configuration.value = cfg;
        })
    }

    async function updateConfig(input: UpdateConfigInput) {
        return await invoke<ReadConfigOutput>('update_user_config', {input: input}).then((cfg: ReadConfigOutput) => {
            configuration.value = cfg
        })
    }

    return {
        configuration,
        getConfig
    }
})