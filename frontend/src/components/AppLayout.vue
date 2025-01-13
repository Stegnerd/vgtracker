<script setup lang="ts"></script>
<template>
  <div class="layout-wrapper flex flex-col">
    <app-topbar />
    <div class="layout-content flex flex-row h-full">
      <app-menu />
      <RouterView />
    </div>
  </div>
</template>

<style>
  .layout-wrapper {
    padding: 1rem;
    height: 100dvh;
  }
</style>
