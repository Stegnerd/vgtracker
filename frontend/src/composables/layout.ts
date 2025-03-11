import { $t, updatePreset, updateSurfacePalette } from "@primeuix/themes";
import Aura from "@primeuix/themes/aura";
import Lara from "@primeuix/themes/lara";
import Nora from "@primeuix/themes/nora";
import { computed, reactive, readonly, ref } from "vue";
import { models } from "../../wailsjs/go/models";
import { PaletteList, SurfaceList } from "../models/theme";

const layoutConfig = reactive({
  preset: "Lara",
  primary: "emerald",
  surface: "zinc",
  darkTheme: true,
  menuMode: "static"
});

const layoutState = reactive({
  staticMenuDesktopInactive: false,
  overlayMenuActive: false,
  profileSidebarVisible: false,
  configSidebarVisible: false,
  staticMenuMobileActive: false,
  menuHoverActive: false,
  activeMenuItem: null
});

const presets = {
  Aura,
  Lara,
  Nora
};
const presetOptions = ref(Object.keys(presets));

const preset = ref(layoutConfig.preset);
const primaryColors = ref(PaletteList);
const surfaces = ref(SurfaceList);

export function useLayout() {
  const setPrimary = (value: string) => {
    layoutConfig.primary = value;
    console.warn("layoutconfig.primary", layoutConfig.primary);
  };

  const setSurface = (value: string) => {
    layoutConfig.surface = value;
  };

  const setPreset = (value: string) => {
    layoutConfig.preset = value;
    console.warn("layoutconfig.preset", layoutConfig.preset);
  };

  const setDefaultValues = (
    primary: models.PaletteColor,
    surface: models.SufaceColor,
    presetStyle: models.PresetConfig,
    isDark: boolean
  ) => {
    layoutConfig.primary = primary;
    layoutConfig.surface = surface;
    layoutConfig.preset = presetStyle;
    layoutConfig.darkTheme = isDark;
    preset.value = layoutConfig.preset;

    updateColors("primary", layoutConfig.primary);
    updateColors("surface", layoutConfig.surface);

    if (!isDark) {
      toggleDarkMode();
    }

    onPresetChange();
  };
  //   const setActiveMenuItem = (item) => {
  //     layoutState.activeMenuItem = item.value || item;
  //   };

  //   const setMenuMode = (mode) => {
  //     layoutConfig.menuMode = mode;
  //   };

  // const toggleDarkMode = () => {
  //   if (!document.startViewTransition) {
  //     executeDarkModeToggle();

  //     return;
  //   }

  //   document.startViewTransition(() => executeDarkModeToggle(event));
  // };

  //   const executeDarkModeToggle = () => {
  //     layoutConfig.darkTheme = !layoutConfig.darkTheme;
  //     document.documentElement.classList.toggle("app-dark");
  //   };

  //   const onMenuToggle = () => {
  //     if (layoutConfig.menuMode === "overlay") {
  //       layoutState.overlayMenuActive = !layoutState.overlayMenuActive;
  //     }

  //     if (window.innerWidth > 991) {
  //       layoutState.staticMenuDesktopInactive =
  //         !layoutState.staticMenuDesktopInactive;
  //     } else {
  //       layoutState.staticMenuMobileActive = !layoutState.staticMenuMobileActive;
  //     }
  //   };

  //   const resetMenu = () => {
  //     layoutState.overlayMenuActive = false;
  //     layoutState.staticMenuMobileActive = false;
  //     layoutState.menuHoverActive = false;
  //   };

  //   const isSidebarActive = computed(
  //     () => layoutState.overlayMenuActive || layoutState.staticMenuMobileActive
  //   );

  const toggleDarkMode = () => {
    const element = document.querySelector("html");
    element!.classList.toggle("app-dark");
    //isDarkMode.value = !isDarkMode.value;
    layoutConfig.darkTheme = !layoutConfig.darkTheme;
  };

  const isDarkTheme = computed(() => layoutConfig.darkTheme);

  //   const getPrimary = computed(() => layoutConfig.primary);

  //   const getSurface = computed(() => layoutConfig.surface);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  function updateColors(type: string, color: any) {
    if (type === "primary") {
      setPrimary(color);
    } else if (type === "surface") {
      setSurface(color);
    }

    applyTheme(type, color);
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  function applyTheme(type: string, color: any) {
    if (type === "primary") {
      updatePreset(getPresetExt());
    } else if (type === "surface") {
      const surfacePalette = surfaces.value.find(
        (s) => s.name === color
      )?.palette;
      updateSurfacePalette(surfacePalette);
    }
  }

  function onPresetChange() {
    console.warn("onPresetChange defualt value", preset.value);
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

  return {
    primaryColors,
    surfaces,
    preset,
    presetOptions,
    layoutConfig: readonly(layoutConfig),
    layoutState: readonly(layoutState),
    //     onMenuToggle,
    //     isSidebarActive,
    isDarkTheme,
    //     getPrimary,
    //     getSurface,
    //     setActiveMenuItem,
    //     toggleDarkMode,
    setDefaultValues,
    toggleDarkMode,
    setPrimary,
    setSurface,
    setPreset,
    updateColors,
    onPresetChange
    //     resetMenu,
    //     setMenuMode
  };
}
