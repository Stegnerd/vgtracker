/// <reference types="vitest/config" />
import { PrimeVueResolver } from "@primevue/auto-import-resolver";
import vue from "@vitejs/plugin-vue";
import { fileURLToPath, URL } from "node:url";
import UnoCSS from "unocss/vite";
import Components from "unplugin-vue-components/vite";
import { defineConfig } from "vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    UnoCSS(),
    Components({
      resolvers: [PrimeVueResolver()]
    })
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url))
    }
  },
  test: {
    globals: true,
    environment: "jsdom"
  }
  // fixes a bug for sass warnings, will be resolved in a higher version of vite
  // https://github.com/vitejs/vite/issues/18164#issuecomment-2365383779
  //   css: {
  //   preprocessorOptions: {
  //     scss: {
  //       api: 'modern-compiler',
  //     },
  //   },
  // },
});
