<script setup lang="ts">
  import { useLayout } from "@/layout/composables/layout";
  import {
    AmberPalette,
    BluePalette,
    CyanPalette,
    EmeraldPalette,
    FuchsiaPalette,
    GraySurfacePalette,
    GreenPalette,
    IndigoPalette,
    LimePalette,
    NeutralSurfacePalette,
    NoirPalette,
    OceanSurfacePalette,
    OrangePalette,
    PaletteContainer,
    PaletteType,
    PinkPalette,
    PurplePalette,
    RosePalette,
    SkyPalette,
    SlateSurfacePalette,
    SohoSurfacePalette,
    StoneSurfacePalette,
    TealPalette,
    ThemePreset,
    ThemePresetType,
    ThemePrimaryColor,
    ThemeSurfaceColor,
    VioletPalette,
    VivaSurfacePalette,
    YellowPalette,
    ZincSurfacePalette
  } from "@/layout/composables/model";
  import { $t, updatePreset, updateSurfacePalette } from "@primevue/themes";
  import { Ref, ref } from "vue";

  const { layoutConfig, setPrimary, setSurface, setPreset, setMenuMode } =
    useLayout();
  //const preset = ref<typeof Aura | typeof Lara>(layoutConfig.preset);
  const preset = ref<ThemePresetType>(layoutConfig.preset);
  const presetObjections = ref(Object.keys(ThemePreset));

  const menuMode = ref(layoutConfig.menuMode);
  const menuModeOptions = ref([
    { label: "Static", value: "static" },
    { label: "Overlay", value: "overlay" }
  ]);

  const primaryColors: Ref<PaletteContainer[]> = ref([
    NoirPalette,
    EmeraldPalette,
    GreenPalette,
    LimePalette,
    OrangePalette,
    AmberPalette,
    YellowPalette,
    TealPalette,
    CyanPalette,
    SkyPalette,
    BluePalette,
    IndigoPalette,
    VioletPalette,
    PurplePalette,
    FuchsiaPalette,
    PinkPalette,
    RosePalette
  ]);

  const surfaces: Ref<PaletteContainer[]> = ref([
    SlateSurfacePalette,
    GraySurfacePalette,
    ZincSurfacePalette,
    NeutralSurfacePalette,
    StoneSurfacePalette,
    SohoSurfacePalette,
    VivaSurfacePalette,
    OceanSurfacePalette
  ]);

  function getPresetExtension() {
    const color = primaryColors.value.find(
      (primary) => primary.name === layoutConfig.primary
    );

    if (color?.name === ThemePrimaryColor.Noir) {
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

  function updateColors(type: "primary" | "surface", color: PaletteContainer) {
    if (type === "primary" && color.type === PaletteType.Primary) {
      setPrimary(color.name as ThemePrimaryColor);
    } else if (type === "surface" && color.type === PaletteType.Surface) {
      setSurface(color.name as ThemeSurfaceColor);
    }

    applyTheme(type, color);
  }

  function applyTheme(type: "primary" | "surface", color: PaletteContainer) {
    if (type === "primary") {
      updatePreset(getPresetExtension());
    } else if (type === "surface") {
      updateSurfacePalette(color.palette);
    }
  }

  function onPresetChange() {
    setPreset(preset.value);
    //const presetValue = ThemePreset[preset.value];
    const surfacePalette = surfaces.value.find(
      (s) => s.name === layoutConfig.surface
    )?.palette;

    $t()
      .preset(preset.value)
      .preset(getPresetExtension())
      .surfacePalette(surfacePalette)
      .use({ useDefaultOptions: true });
  }

  function onMenuModeChange() {
    setMenuMode(menuMode.value);
  }
</script>

<template></template>

<style lang="scss" scoped></style>
