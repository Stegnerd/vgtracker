<script setup lang="ts">
import NavBar from '../components/NavBar.vue';
</script>

<template>
  <div class="min-h-screen flex flex-col bg-gray-300 dark:bg-slate-800 ">
    <div class='relative flex flex-grow'>
      <NavBar class='bg-white shadow-sm p-6 space-y-6 w-64'></NavBar>
      <main class='bg-gray-100 flex-1 p-6 main-card'>
        <RouterView></RouterView>
      </main>
    </div>
  </div>
</template>

<style scoped>
.main-card {
  margin: 2rem;
  border-radius: 10px;
}
</style>
