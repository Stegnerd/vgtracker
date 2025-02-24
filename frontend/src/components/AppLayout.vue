<script setup lang="ts"></script>
<template>
  <div class="layout-wrapper flex">
    <app-menu class="pt-menu" />
    <div class="layout-content flex flex-col h-full w-full">
      <app-topbar />
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
