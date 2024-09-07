<script setup lang="ts"></script>
<template>
  <div class="layout-wrapper flex flex-col">
    <app-topbar></app-topbar>
    <div class="layout-content">
      <RouterView />
    </div>
  </div>
</template>

<style>
  .layout-wrapper {
    background-color: rebeccapurple;
    height: 98vh;
  }
</style>
