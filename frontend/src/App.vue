<script lang="ts" setup>
	import DrawerComponent from './components/Drawer.vue';
</script>

<template>
	<v-app>
		<v-navigation-drawer expand-on-hover rail permanent>
			<DrawerComponent />
		</v-navigation-drawer>
		<v-main>
			<router-view></router-view>
		</v-main>
	</v-app>
</template>

<style></style>
