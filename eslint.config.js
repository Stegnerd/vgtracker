import typescriptParser from "@typescript-eslint/parser";
import eslintConfigPrettier from "eslint-config-prettier";
import importPlugin from "eslint-plugin-import";
import simpleImportSort from "eslint-plugin-simple-import-sort";
import pluginVue from "eslint-plugin-vue";
import parser from "vue-eslint-parser";


export default [
    // add more generic rulesets here, such as:
    // js.configs.recommended,
    ...pluginVue.configs['flat/recommended'],
    // ...pluginVue.configs['flat/vue2-recommended'], // Use this if you are using Vue.js 2.x.
    {
        plugins: {
            "simple-import-sort": simpleImportSort,
            "import": importPlugin
        },
        files: ["src/**/*.ts", "src/**/*.vue"],
        rules: {
            // override/add rules settings here, such as:
            // 'vue/no-unused-vars': 'error'
            "semi": ["error", "always"],
            "simple-import-sort/imports": "error",
            "simple-import-sort/exports": "error",
            "import/first": "error",
            "import/newline-after-import": "error",
            "import/no-duplicates": "error"
        },
        languageOptions: {
            parser: parser,
            parserOptions: {
                parser: typescriptParser,
                sourceType: "module"
            }
        }
    },
    eslintConfigPrettier
];
