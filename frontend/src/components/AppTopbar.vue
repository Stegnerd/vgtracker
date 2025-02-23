<script setup lang="ts">
import { AutoCompleteCompleteEvent, AutoCompleteOptionSelectEvent } from "primevue";
import { useDialog } from "primevue/usedialog";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { SearchMainGames } from "../../wailsjs/go/controllers/IgdbController";
import { igdb } from "../../wailsjs/go/models";
import { useIGDBSelection } from "../composables/igdbSelection";
import { useLayout } from "../composables/layout";
import AppThemeConfigurator from "./AppThemeConfigurator.vue";

const router = useRouter();
const { setGameSelection } = useIGDBSelection();
const {toggleDarkMode} = useLayout();
const dialog = useDialog();
const isDarkMode = ref(true);

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
const search = (event: AutoCompleteCompleteEvent) => {

       console.warn('EVENT', event)
    if (event.query !== "") {
      SearchMainGames(event.query).then((results) => {
        console.warn('search results', results)
        filterResults.value = results.items
      })
    }
}

const selected = (event: AutoCompleteOptionSelectEvent) => {
  searchInput.value = ''
  const result = event.value as igdb.VGTGame
  setGameSelection(result)
  filterResults.value = [];
  router.push({path:'/game-details'})
}

</script>
<template>
  <div class="topbar-container flex flex-row justify-between items-center pb-10">
    <div />
    <div class="flex justify-center">
      <InputGroup class="min-w-lg">
        <InputGroupAddon>
          <i class="i-mdi-magnify" />
        </InputGroupAddon>
        <!-- option-label="name" -->
        <AutoComplete
          v-model="searchInput"
          spellcheck="false"
          input-id="daltonfick"
          :suggestions="filterResults"
          :delay="1000"
          @complete="search"
          @option-select="selected"
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
