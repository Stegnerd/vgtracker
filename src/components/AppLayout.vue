<script setup lang="ts"></script>
<template>
  <div class="layout-wrapper flex flex-col">
    <app-topbar></app-topbar>
    <div class="layout-content flex flex-row h-full">
      <app-menu></app-menu>
      <RouterView />
    </div>
  </div>
</template>

<style>
  .layout-wrapper {
    padding: 8px;
    height: 100dvh;
  }
</style>
