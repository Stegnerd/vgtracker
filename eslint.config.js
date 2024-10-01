import vueTsEslintConfig from "@vue/eslint-config-typescript";
import pluginVue from 'eslint-plugin-vue';

export default [
  // add more generic rulesets here, such as:
  // js.configs.recommended,
  ...pluginVue.configs['flat/recommended'],
  ...vueTsEslintConfig(),
  {
    // https://github.com/eslint/eslint/discussions/18304#discussioncomment-9069706
    ignores: ["src/vite-env.d.ts"],
  },
  // ...pluginVue.configs['flat/vue2-recommended'], // Use this if you are using Vue.js 2.x.
  {
    rules: {
      // override/add rules settings here, such as:
      // 'vue/no-unused-vars': 'error'
    }
  }
];
