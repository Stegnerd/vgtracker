<script setup lang="ts">
import { ref } from "vue";

const items = ref([
  {
    items: [
      {
        label: 'Search',
        icon: 'pi pi-search',
        link: '/search',
        shortcut: '⌘+O'
      },
      {
        label: 'Settings',
        icon: 'pi pi-cog',
        link: '/settings',
        shortcut: '⌘+O'
      }
    ]
  }
]);
</script>

<template>
  <div class="card ">
    <Menu :model="items" class="min-h-full">
      <template #item="{ item, props }">
        <RouterLink :to="item.link" v-ripple class="flex align-items-center" active-class="active-menu"
          v-bind="props.action">
          <span :class="item.icon" />
          <span class="ml-2">{{ item.label }}</span>
          <!-- <Badge v-if="item.badge" class="ml-auto" :value="item.badge" /> -->
          <!-- <span v-if="item.shortcut" class="ml-auto border-1 surface-border border-round surface-100 text-xs p-1">{{
          item.shortcut }}</span> -->
        </RouterLink>
      </template>
    </Menu>
  </div>
</template>

<style scoped>
.card {
  padding: 2rem;
  border-radius: 10px;
}

.active-menu {
  /* background-color: theme('colors.primary-500'); */
}
</style>
