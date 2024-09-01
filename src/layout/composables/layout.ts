import { reactive, readonly } from "vue";
import {
  LayoutConfig,
  ThemeMenuMode,
  ThemePreset,
  ThemePresetType,
  ThemePrimaryColor,
  ThemeSurfaceColor
} from "./model";

const layoutConfig: LayoutConfig = reactive({
  preset: ThemePreset.Lara,
  primary: ThemePrimaryColor.Emerald,
  surface: ThemeSurfaceColor.Slate,
  darkTheme: false,
  menuMode: ThemeMenuMode.Static
});

export function useLayout() {
  const setPrimary = (value: ThemePrimaryColor) => {
    layoutConfig.primary = value;
  };

  const setSurface = (value: ThemeSurfaceColor) => {
    layoutConfig.surface = value;
  };

  //const setPreset = (value: typeof ThemePreset) => {
  const setPreset = (value: ThemePresetType) => {
    layoutConfig.preset = value;
  };

  const setMenuMode = (mode: ThemeMenuMode) => {
    layoutConfig.menuMode = mode;
  };

  const toggleDarkMode = () => {
    if (!document.startViewTransition) {
      executeDarkModeToggle();

      return;
    }

    document.startViewTransition(() => executeDarkModeToggle());
  };

  const executeDarkModeToggle = () => {
    layoutConfig.darkTheme = !layoutConfig.darkTheme;
    document.documentElement.classList.toggle("app-dark");
  };

  return {
    layoutConfig: readonly(layoutConfig),
    setPrimary,
    setSurface,
    setPreset,
    setMenuMode,
    toggleDarkMode
  };
}
