import {
  defineConfig,
  presetAttributify,
  presetIcons,
  presetUno
} from "unocss";

export default defineConfig({
  //rules: [["m-1", { margin: "1px" }]]
  presets: [presetUno(), presetAttributify(), presetIcons()]
});
