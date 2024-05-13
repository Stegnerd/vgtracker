<script setup lang="ts">
  function onSubmit() {
    console.warn("here1");
  }
</script>

<template>
  SearchView
  <PrimeButton
    label="Submit"
    type="submit"
    data-testid="submit-button"
    @click="onSubmit()"
  />
</template>
