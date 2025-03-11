<script setup lang="ts">
  import { useLayout } from "../composables/layout";
  const { primaryColors, surfaces,preset, presetOptions,layoutConfig, onPresetChange, updateColors, isDarkTheme } = useLayout();



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
          @click="updateColors('primary', primaryColor.name)"
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
          @click="updateColors('surface', surface.name)"
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
          @change="onPresetChange()"
        />
      </div>
    </div>
  </div>
</template>
