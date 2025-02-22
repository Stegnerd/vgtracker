import { createLocalFontProcessor } from "@unocss/preset-web-fonts/local";
import {
  defineConfig,
  presetAttributify,
  presetIcons,
  presetUno,
  presetWebFonts
} from "unocss";

export default defineConfig({
  //rules: [["m-1", { margin: "1px" }]]
  presets: [
    presetUno(),
    presetAttributify(),
    presetIcons(),
    presetWebFonts({
      provider: "bunny",
      fonts: {
        sans: "Roboto",
        mono: ["Fira Code", "Fira Mono:400,700"]
      },
      // This will download the fonts and serve them locally
      processors: createLocalFontProcessor({
        // Directory to cache the fonts
        cacheDir: "node_modules/.cache/unocss/fonts",

        // Directory to save the fonts assets
        fontAssetsDir: "public/assets/fonts",

        // Base URL to serve the fonts from the client
        fontServeBaseUrl: "/assets/fonts"
      })
    })
  ],
  rules: [["pt-menu", { "padding-top": "5.2rem" }]]
});
