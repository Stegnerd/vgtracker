import { createRouter, createWebHistory, Router, RouteRecordRaw } from 'vue-router';
import Gallery from './views/Gallery.vue';
import Profile from './views/Profile.vue';
import MyGames from './views/MyGames.vue';

const routes: RouteRecordRaw[] = [
	{ path: '/', component: Gallery },
	{ path: '/gallery', component: Gallery },
	{ path: '/profile', component: Profile },
	{ path: '/my-games', component: MyGames }
];

const router: Router = createRouter({
	history: createWebHistory(),
	routes
});

export default router;
