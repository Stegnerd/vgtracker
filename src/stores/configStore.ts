import {defineStore} from "pinia";
import {invoke} from "@tauri-apps/api/tauri";
import {Config} from '../../src-tauri/frontend_models/config.ts'
import {ref} from "vue";

export const useConfigStore = defineStore('config', () => {
    const configuration = ref<Config | undefined>();

    async function getConfig() {
        return await invoke<Config>('get_user_config').then((cfg: Config) => {
            configuration.value = cfg;
        })
    }

    return {
        configuration,
        getConfig
    }
})