import AppLayout from "@/components/AppLayout.vue";
import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  linkActiveClass: "active-menu-item",
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: AppLayout,
      children: [
        {
          path: "/",
          redirect: "/settings"
        },
        {
          path: "/settings",
          name: "settings",
          component: () => import("@/pages/Settings.vue")
        },
        {
          path: "/dashboard",
          name: "dashboard",
          component: () => import("@/pages/Dashboard.vue")
        }
      ]
    }
  ]
});
export default router;
