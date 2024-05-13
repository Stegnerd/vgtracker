<script setup lang="ts">
  import { ref } from "vue";

  import ThemeSwitcher from "@/components/ThemeSwitcher.vue";

  const items = ref([
    {
      items: [
        {
          label: "Search",
          icon: "pi pi-search",
          link: "/search",
          shortcut: "⌘+O"
        },
        {
          label: "Profile",
          icon: "pi pi-user",
          link: "/profile",
          shortcut: "⌘+O"
        },
        {
          label: "Settings",
          icon: "pi pi-cog",
          link: "/settings",
          shortcut: "⌘+O"
        }
      ]
    }
  ]);
</script>

<template>
  <div class="p-8 rounded-[10px]">
    <PrimeMenu :model="items" class="min-h-full pb-0">
      <template #item="{ item, props }">
        <RouterLink
          v-ripple
          :to="item.link"
          class="flex align-items-center"
          active-class="active-menu"
          v-bind="props.action"
        >
          <span :class="item.icon" />
          <span class="ml-2">{{ item.label }}</span>
          <!-- <Badge v-if="item.badge" class="ml-auto" :value="item.badge" /> -->
          <!-- <span v-if="item.shortcut" class="ml-auto border-1 surface-border border-round surface-100 text-xs p-1">{{
          item.shortcut }}</span> -->
        </RouterLink>
      </template>
      <template #end>
        <div class="flex flex-col justify-end h-[calc(100vh-242px)]">
          <ThemeSwitcher />
        </div>
      </template>
    </PrimeMenu>
  </div>
</template>

<style scoped></style>
