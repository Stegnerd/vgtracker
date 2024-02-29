import { RouteRecordRaw, createRouter, createWebHashHistory } from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    redirect: "/search"
  },
  {
    meta: {
      title: "Search"
    },
    path: "/search",
    name: "search",
    component: () => import("@/views/SearchView.vue")
  },
  {
    meta: {
      title: "Settings"
    },
    path: "/settings",
    name: "settings",
    component: () => import("@/views/SettingView.vue")
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
