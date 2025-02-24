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
          path: "/game-details",
          name: "game-details",
          component: () => import("@/pages/AppGameDetail.vue")
        },
        {
          path: "/settings",
          name: "settings",
          component: () => import("@/pages/AppSettings.vue")
        },
        {
          path: "/dashboard",
          name: "dashboard",
          component: () => import("@/pages/AppDashboard.vue")
        }
      ]
    }
  ]
});
export default router;
