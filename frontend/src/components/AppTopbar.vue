<script setup lang="ts">
import {watchDebounced} from "@vueuse/core";
import {useDialog} from "primevue/usedialog";
import {ref} from "vue";
import {Search} from "../../wailsjs/go/controllers/IgdbController";
import {igdb} from "../../wailsjs/go/models";
import {useLayout} from "../composables/layout";
import AppThemeConfigurator from "./AppThemeConfigurator.vue";

const {toggleDarkMode} = useLayout();

const dialog = useDialog();

const isDarkMode = ref(true);

// function toggleDarkMode() {
//   const element = document.querySelector("html");
//   element!.classList.toggle("app-dark");
//   isDarkMode.value = !isDarkMode.value;
// }

function openThemeDialog() {
  dialog.open(AppThemeConfigurator, {
    props: {
      modal: false,
      position: "topright"
    }
  });
}

const searchInput = ref('');
const filterResults = ref<igdb.VGTGame[]>([]);
watchDebounced(searchInput, async (newSearch) => {
      console.warn('NEW SEARCH', newSearch)
      Search(newSearch).then((results) => {
        console.warn('search results', results)
        filterResults.value = results.items
      })
    },
    {debounce: 500, maxWait: 1000}
)
</script>
<template>
  <div class="topbar-container flex flex-row justify-between items-center pb-10">
    <div></div>
    <div class="flex justify-center">
      <InputGroup class="min-w-lg">
        <InputGroupAddon>
          <i class="i-mdi-magnify"/>
        </InputGroupAddon>
        <!-- <InputText
          v-model="searchInput"
          placeholder="Search"
        /> -->
        <AutoComplete
            v-model="searchInput"
            :suggestions="filterResults"
        >
          <template #option="slotProps">
            <div>
              <p>{{ slotProps.option.name }}</p>
            </div>
          </template>
        </AutoComplete>
      </InputGroup>
    </div>
    <div class="flex gap-6">
      <Button
          :icon="isDarkMode ? 'i-mdi-weather-night' : 'i-mdi-white-balance-sunny'"
          @click="toggleDarkMode()"
      />
      <Button
          icon="i-mdi-palette"
          @click="openThemeDialog()"
      />
    </div>
  </div>
</template>
<style></style>
