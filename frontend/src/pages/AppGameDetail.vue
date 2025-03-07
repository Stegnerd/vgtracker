<script setup lang="ts">
import { onMounted } from 'vue';
import { GetGameDetails } from "../../wailsjs/go/controllers/GameDetailController";
import { useIGDBSelection } from '../composables/igdbSelection';


const { gameSelection } = useIGDBSelection()


onMounted(async () => {
  GetGameDetails(gameSelection.value?.id || 0).then( (result) => {
    console.warn('detail list', result)
  })
})


</script>

<template>
  <p ref="testing">
    SELECTED GAMES: {{ gameSelection?.name }}
  </p>
</template>
<style></style>
