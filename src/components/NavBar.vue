<script setup lang="ts">
import { ref } from "vue";

const items = ref([
  {
    items: [
      {
        label: 'Search',
        icon: 'pi pi-search',
        link: 'search',
        shortcut: '⌘+O'
      },
      {
        label: 'Settings',
        icon: 'pi pi-cog',
        link: 'setting',
        shortcut: '⌘+O'
      },
      {
        label: 'Logout',
        icon: 'pi pi-sign-out',
        shortcut: '⌘+Q'
      }
    ]
  }
]);
</script>

<template>
  <div class="card ">
    <Menu :model="items" class="min-h-full">
      <template #item="{ item, props }">
        <a v-ripple class="flex align-items-center" v-bind="props.action">
          <span :class="item.icon" />
          <span class="ml-2">{{ item.label }}</span>
          <!-- <Badge v-if="item.badge" class="ml-auto" :value="item.badge" /> -->
          <!-- <span v-if="item.shortcut" class="ml-auto border-1 surface-border border-round surface-100 text-xs p-1">{{
          item.shortcut }}</span> -->
        </a>
      </template>
    </Menu>
  </div>
</template>

<style scoped>
.card {
  background: var(--surface-card);
  padding: 2rem;
  border-radius: 10px;
}
</style>
