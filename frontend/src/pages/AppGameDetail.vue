<script setup lang="ts">
import { onMounted } from 'vue';
import { GetGameDetails } from "../../wailsjs/go/controllers/GameDetailController";
import AppCard from '../components/AppCard.vue';
import { useIGDBSelection } from '../composables/igdbSelection';


const { gameSelection } = useIGDBSelection()
console.warn('here is the selected Game', gameSelection.value)
// gameSelection.value = {
// id: 82373,
// coverURL : "https://images.igdb.com/igdb/image/upload/t_cover_big/co2xcl.jpg",
// firstReleaseYear: 2018,
// gameType: 'Remaster',
// genres: ["Adventure", "Role-Playing (RPG)", "Simulator", "Strategy", "Tactical", "Turn-based strategy (TBS)"],
// name : "Disgaea 1 Complete",
// platforms: ["Android", "Nintendo Switch", "Playstation 4", "iOS"],
// thumbnailURL: 'https://images.igdb.com/igdb/image/upload/t_thumb/co2xcl.jpg'
// } as igdb.VGTGame


onMounted(async () => {
  GetGameDetails(gameSelection.value?.id || 0).then( (result) => {
    console.warn('detail list', result)
  })
})

// TODO: SET THIS UP TO USE PRIMEVUE FORMS INSTEAD?

</script>

<template>
  <AppCard>
    <div class="flex">
      <Image
        :src="gameSelection?.coverURL"
        class="pr-4"
      />
      <div class="flex flex-col">
        <span class="text-3xl">
          {{ gameSelection?.name }}
        </span>
        <p>testin</p>
      </div>
    </div>
    <Divider />
    <div class="flex">
      <p>{{ gameSelection?.summary }}</p>
    </div>
  </AppCard>
</template>
<style></style>
