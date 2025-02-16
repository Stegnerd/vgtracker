// import js from "@eslint/js";
// import eslintPluginVue from "eslint-plugin-vue";
// import ts from "typescript-eslint";

// export default ts.config(
//   js.configs.recommended,
//   ...ts.configs.recommended,
//   ...eslintPluginVue.configs["flat/recommended"],
//   {
//     ignores: ["**/wailsjs/*", "src/vite-env.d.ts", "node_modules/*", "dist/*"]
//   },
//   {
//     files: ["*.vue", "**/*.vue", "**/*.ts"],

//     languageOptions: {
//       parserOptions: {
//         parser: "@typescript-eslint/parser"
//       }
//     }
//   }
// );

import pluginVitest from "@vitest/eslint-plugin";
import {
  defineConfigWithVueTs,
  vueTsConfigs
} from "@vue/eslint-config-typescript";
import pluginVue from "eslint-plugin-vue";

// To allow more languages other than `ts` in `.vue` files, uncomment the following lines:
// import { configureVueProject } from '@vue/eslint-config-typescript'
// configureVueProject({ scriptLangs: ['ts', 'tsx'] })
// More info at https://github.com/vuejs/eslint-config-typescript/#advanced-setup

export default defineConfigWithVueTs(
  {
    name: "app/files-to-lint",
    files: ["**/*.{ts,vue}"]
  },

  {
    name: "app/files-to-ignore",
    ignores: [
      "**/dist/**",
      "**/dist-ssr/**",
      "**/wailsjs/*",
      "node_modules/*",
      "src/vite-env.d.ts"
    ]
  },

  pluginVue.configs["flat/recommended"],
  vueTsConfigs.recommended,
  {
    rules: {
      "vue/max-attributes-per-line": [
        "error",
        {
          singleline: {
            max: 1
          },
          multiline: {
            max: 1
          }
        }
      ]
    }
  },

  {
    ...pluginVitest.configs.recommended,
    files: ["src/**/__tests__/*"]
  }
);
