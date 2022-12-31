<script lang="ts" setup>
	import DrawerComponent from './components/DrawerComponent.vue';
</script>

<template>
	<v-app>
		<v-navigation-drawer expand-on-hover rail permanent>
			<DrawerComponent />
		</v-navigation-drawer>
		<v-main>
			<!--			<img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png" />-->
			<!--			<HelloWorld />-->
		</v-main>
	</v-app>
</template>

<style>
	#logo {
		display: block;
		width: 50%;
		height: 50%;
		margin: auto;
		padding: 10% 0 0;
		background-position: center;
		background-repeat: no-repeat;
		background-size: 100% 100%;
		background-origin: content-box;
	}
</style>
