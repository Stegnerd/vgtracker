<script setup lang="ts">
import { useIGDBSelection } from '../composables/igdbSelection';


const { gameSelection } = useIGDBSelection()

</script>

<template>
  <p>SELECTED GAMES: {{ gameSelection?.name }}</p>
</template>
<style></style>
