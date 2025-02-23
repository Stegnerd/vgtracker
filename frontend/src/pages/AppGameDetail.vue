<script setup lang="ts">
import { onMounted, useTemplateRef } from 'vue';
import { useIGDBSelection } from '../composables/igdbSelection';


const { gameSelection } = useIGDBSelection()
const testRef = useTemplateRef('testing');

onMounted(() => {
  testRef.value?.focus();
})


</script>

<template>
  <p ref="testing">
    SELECTED GAMES: {{ gameSelection?.name }}
  </p>
</template>
<style></style>
