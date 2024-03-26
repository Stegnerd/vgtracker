<script setup lang="ts">
import {useConfigStore} from "@/stores/configStore.ts";
import {storeToRefs} from "pinia";
import {ref} from "vue";

const configStore = useConfigStore();
const {configuration} = storeToRefs(configStore)


//const iconClass = ref('pi-moon');
console.warn('toggler conf', configuration.value)
let iconClass = ref(configuration.value?.theme === 'Light' ? 'pi-sun' : 'pi-moon')

// iconClass.value = configuration.value?.theme === 'Light' ? 'pi-sun' : 'pi-moon'

async function onThemeToggle() {
  const root = document.getElementsByTagName('html')[0];
  root.classList.toggle('dark');
  const input = configuration.value?.theme === 'Light' ? 'Dark' : 'Light'
  await configStore.updateTheme(input).then(() => {
    console.warn('component await', configuration.value)

    iconClass.value = configuration.value?.theme === 'Light' ? 'pi-sun' : 'pi-moon';
    console.warn('iconClass.value', iconClass.value)
  })
}

</script>

<template>
  <div class="card flex justify-end p-2 mb-4">
    <ul class="flex list-none m-0 p-0 gap-2 items-center">
      <li>
        <button
            type="button"
            class="inline-flex border w-8 h-8 p-0 items-center justify-center surface-0 dark:surface-800 border border-surface-200 dark:border-surface-600 rounded"
            @click="onThemeToggle"
        >
          <i :class="`dark:text-white pi ${iconClass}`"/>
        </button>
      </li>
    </ul>
  </div>
</template>

<style scoped>

</style>