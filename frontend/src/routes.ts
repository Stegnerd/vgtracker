import AppLayout from "@/components/AppLayout.vue";
import { createRouter, createWebHistory } from "vue-router";
import AppBacklog from "./pages/AppBacklog.vue";
import AppGameDetail from "./pages/AppGameDetail.vue";

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
          redirect: "/dashboard"
        },
        {
          path: "/dashboard",
          name: "dashboard",
          component: () => import("@/pages/AppDashboard.vue")
        },
        {
          path: "/library",
          component: () => import("@/pages/AppLibrary.vue"),
          children: [
            {
              path: "backlog",
              name: "backlog",
              component: AppBacklog
            },
            {
              path: "game-details",
              name: "game-details",
              component: AppGameDetail
            }
          ]
        },
        {
          path: "/settings",
          name: "settings",
          component: () => import("@/pages/AppSettings.vue")
        }
      ]
    }
  ]
});
export default router;
