<script setup lang="ts">
  import { MenuItem } from "primevue/menuitem";
  import { ref } from "vue";

  // https://github.com/primefaces/primevue/issues/6103
  // console error for aria hidden
  const menuItems = ref<MenuItem[]>([
    {
      separator: true
    },
    {
      label: "Dashboard",
      icon: "i-mdi-home",
      route: "/dashboard"
    },
    {
      label: "Settings",
      icon: "i-mdi-cog",
      route: "/settings"
    }
  ]);
</script>

<template>
  <div class="menu-item pl-2 pr-2">
    <Menu :model="menuItems" class="menu">
      <template #start>
        <span class="inline-flex items-center gap-1 px-2 py-2">
          <span class="i-mdi-nintendo-game-boy" />
          <span class="text-l font-semibold">VGTracker</span>
        </span>
      </template>
      <template #item="{ item, props }">
        <router-link
          active-class="active-menu-item"
          v-if="item.route"
          :to="item.route"
          v-slot="{ href, navigate }"
        >
          <a v-ripple :href="href" v-bind="props.action" @click="navigate">
            <span :class="item.icon" />
            <span>{{ item.label }}</span>
          </a>
        </router-link>
      </template>
      <template #end></template>
    </Menu>
  </div>
</template>

<style lang="scss">
  // .menu {
  //   height: 100%;
  // }

  .menu-item {
    a {
      text-decoration: none;
      color: var(--p-text-color);
    }
  }

  .active-menu-item {
    a {
      color: var(--primary-color);
    }
  }
</style>
