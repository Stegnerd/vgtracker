<script setup lang="ts">
  const t = false;
</script>
<template>
  Dashboard
  <p v-if="t" />
  <Skeleton
    v-else
    shape="rectangle"
    width="90px"
    height="90px"
    animation="none"
    class="thumbnail-skeleton"
  />
</template>
<style>
.thumbnail-skeleton {
  width: 90px;
  height: 90px;
}

</style>
