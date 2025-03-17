<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetGameDetails, UpsertGameDetails } from "../../wailsjs/go/controllers/GameDetailController";
import { gamedetails } from '../../wailsjs/go/models';
import AppCard from '../components/AppCard.vue';
import { useIGDBSelection } from '../composables/igdbSelection';

type updateType = 'wishlist' | 'owned' | 'beaten'
const { gameSelection } = useIGDBSelection()
const backlogDetails = ref<gamedetails.GetGameDetailOutput>(new gamedetails.GetGameDetailOutput())

onMounted(async () => {
  GetGameDetails(gameSelection.value?.id || 0).then( (result) => {
    backlogDetails.value = result;
  })
})

const updateBacklogStats = (selectedAction: updateType) => {
  const input = {
    id: backlogDetails.value.id === 0 ? undefined : backlogDetails.value.id,
    igdbID: gameSelection.value?.id,
    isWishlisted: selectedAction === 'wishlist' ? !backlogDetails.value.isWishlisted : backlogDetails.value.isWishlisted,
    isOwned: selectedAction === 'owned' ? !backlogDetails.value.isOwned : backlogDetails.value.isOwned,
    isBeaten: selectedAction === 'beaten' ? !backlogDetails.value.isBeaten : backlogDetails.value.isBeaten,
    name: gameSelection.value?.name,
    thumbnailURL: gameSelection.value?.thumbnailURL,
    coverURL: gameSelection.value?.coverURL
  } as gamedetails.UpsertGameDetailInput

  UpsertGameDetails(input).then((result) => {
    backlogDetails.value = result
  })
}

</script>

<template>
  <AppCard>
    <div class="flex">
      <Image
        :src="gameSelection?.coverURL"
        class="pr-4"
      />
      <div class="flex flex-col flex-justify-between">
        <div>
          <span class="text-3xl">
            {{ gameSelection?.name }}
          </span>
          <p>testin</p>
        </div>
        <div class="button-container flex ">
          <Button
            label="Wish List"
            icon-pos="top"
            class="mr-4"
            :severity="backlogDetails.isWishlisted ? '':'secondary'"
            :icon="backlogDetails.isWishlisted ? 'i-mdi-star' : 'i-mdi-star-outline'"
            @click="updateBacklogStats('wishlist')"
          />
          <Button
            label="Owned"
            icon-pos="top"
            class="mr-4"
            :severity="backlogDetails.isOwned ? '':'contrast'"
            :icon="backlogDetails.isOwned ? 'i-mdi-check-bold' : 'i-mdi-check-outline' "
            @click="updateBacklogStats('owned')"
          />
          <Button
            label="Beat"
            icon-pos="top"
            :severity="backlogDetails.isBeaten ? '':'contrast'"
            :icon="backlogDetails.isBeaten ? 'i-mdi-check-bold' : 'i-mdi-check-outline' "
            @click="updateBacklogStats('beaten')"
          />
        </div>
      </div>
    </div>
    <Divider />
    <div class="flex">
      <p>{{ gameSelection?.summary }}</p>
    </div>
  </AppCard>
</template>
<style></style>
