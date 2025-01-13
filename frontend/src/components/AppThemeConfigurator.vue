<script setup lang="ts">
  import { $t, updatePreset, updateSurfacePalette } from "@primevue/themes";
  import Aura from "@primevue/themes/aura";
  import Lara from "@primevue/themes/lara";
  import { ref } from "vue";
  import { useLayout } from "../composables/layout";
  import { PaletteList, SurfaceList } from "../models/theme";

  const { layoutConfig, setPrimary, setSurface, setPreset, isDarkTheme } =
    useLayout();

  const presets = {
    Aura,
    Lara
  };
  const preset = ref(layoutConfig.preset);
  const presetOptions = ref(Object.keys(presets));
  const primaryColors = ref(PaletteList);
  const surfaces = ref(SurfaceList);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  function updateColors(type: string, color: any) {
    if (type === "primary") {
      setPrimary(color.name);
    } else if (type === "surface") {
      setSurface(color.name);
    }

    applyTheme(type, color);
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  function applyTheme(type: string, color: any) {
    if (type === "primary") {
      updatePreset(getPresetExt());
    } else if (type === "surface") {
      updateSurfacePalette(color.palette);
    }
  }

  function onPresetChange() {
    setPreset(preset.value);

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const presetValue = (presets as any)[preset.value];
    const surfacePalette = surfaces.value.find(
      (s) => s.name === layoutConfig.surface
    )?.palette;

    $t()
      .preset(presetValue)
      .preset(getPresetExt())
      .surfacePalette(surfacePalette)
      .use({ useDefaultOptions: true });
  }

  function getPresetExt() {
    const color = primaryColors.value.find(
      (c) => c.name === layoutConfig.primary
    );

    if (color?.name === "noir") {
      return {
        semantic: {
          primary: {
            50: "{surface.50}",
            100: "{surface.100}",
            200: "{surface.200}",
            300: "{surface.300}",
            400: "{surface.400}",
            500: "{surface.500}",
            600: "{surface.600}",
            700: "{surface.700}",
            800: "{surface.800}",
            900: "{surface.900}",
            950: "{surface.950}"
          },
          colorScheme: {
            light: {
              primary: {
                color: "{primary.950}",
                contrastColor: "#ffffff",
                hoverColor: "{primary.800}",
                activeColor: "{primary.700}"
              },
              highlight: {
                background: "{primary.950}",
                focusBackground: "{primary.700}",
                color: "#ffffff",
                focusColor: "#ffffff"
              }
            },
            dark: {
              primary: {
                color: "{primary.50}",
                contrastColor: "{primary.950}",
                hoverColor: "{primary.200}",
                activeColor: "{primary.300}"
              },
              highlight: {
                background: "{primary.50}",
                focusBackground: "{primary.300}",
                color: "{primary.950}",
                focusColor: "{primary.950}"
              }
            }
          }
        }
      };
    } else {
      return {
        semantic: {
          primary: color?.palette,
          colorScheme: {
            light: {
              primary: {
                color: "{primary.500}",
                contrastColor: "#ffffff",
                hoverColor: "{primary.600}",
                activeColor: "{primary.700}"
              },
              highlight: {
                background: "{primary.50}",
                focusBackground: "{primary.100}",
                color: "{primary.700}",
                focusColor: "{primary.800}"
              }
            },
            dark: {
              primary: {
                color: "{primary.400}",
                contrastColor: "{surface.900}",
                hoverColor: "{primary.300}",
                activeColor: "{primary.200}"
              },
              highlight: {
                background:
                  "color-mix(in srgb, {primary.400}, transparent 84%)",
                focusBackground:
                  "color-mix(in srgb, {primary.400}, transparent 76%)",
                color: "rgba(255,255,255,.87)",
                focusColor: "rgba(255,255,255,.87)"
              }
            }
          }
        }
      };
    }
  }
</script>
<template>
  <div class="flex flex-col w-64">
    <div>
      <span class="text-sm text-mut">Primary</span>
      <div class="pt-2 flex gap-2 flex-wrap justify-between">
        <button
          v-for="primaryColor of primaryColors"
          :key="primaryColor.name"
          type="button"
          :title="primaryColor.name"
          :class="[
            'border-none w-5 h-5 rounded-full p-0 cursor-pointer outline-none outline-offset-1',
            { 'outline-primary': layoutConfig.primary === primaryColor.name }
          ]"
          :style="{
            backgroundColor: `${
              primaryColor.name === 'noir'
                ? 'var(--p-text-color)'
                : primaryColor.palette['500']
            }`
          }"
          @click="updateColors('primary', primaryColor)"
        />
      </div>
    </div>
    <div class="pt-2">
      <span>Surface</span>
      <div class="pt-2 flex gap-2 flex-wrap justify-between">
        <button
          v-for="surface of surfaces"
          :key="surface.name"
          type="button"
          :title="surface.name"
          :class="[
            'border-none w-5 h-5 rounded-full p-0 cursor-pointer outline-none outline-offset-1',
            {
              'outline-primary': layoutConfig.surface
                ? layoutConfig.surface === surface.name
                : isDarkTheme
                ? surface.name === 'zinc'
                : surface.name === 'slate'
            }
          ]"
          :style="{ backgroundColor: `${surface.palette['500']}` }"
          @click="updateColors('surface', surface)"
        />
      </div>
    </div>
    <div class="flex flex-col pt-2">
      <span>Presets</span>
      <div class="pt-2">
        <SelectButton
          v-model="preset"
          :options="presetOptions"
          :allow-empty="false"
          @change="onPresetChange"
        />
      </div>
    </div>
  </div>
</template>
