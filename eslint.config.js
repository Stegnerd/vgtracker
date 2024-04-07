import typescriptParser from "@typescript-eslint/parser";
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
        },
        files: ["src/**/*.ts", "src/**/*.vue"],
        rules: {
            // override/add rules settings here, such as:
            // 'vue/no-unused-vars': 'error'
            "semi": ["error", "always"],
            //"sort-imports": ["error", {
                    //   "ignoreCase": true, 
                    //     "ignoreDeclarationSort": true 
                // "ignoreCase": false,
                // "ignoreDeclarationSort": false,
                // "ignoreMemberSort": false,
                // "memberSyntaxSortOrder": ["none", "all", "multiple", "single"],
                // "allowSeparatedGroups": false
            //}],
            "simple-import-sort/imports": "error",
            "simple-import-sort/exports": "error",

        },
        languageOptions: {
            parser: parser,
            parserOptions: {
                parser: typescriptParser,
                sourceType: "module"
            }
        }
    }
];
