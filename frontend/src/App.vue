<script lang="ts" setup>
	import HelloWorld from './components/HelloWorld.vue';
</script>

<template>
	<v-app>
		<v-navigation-drawer expand-on-hover rail permanent>
			<v-list>
				<v-list-item
					prepend-avatar="https://randomuser.me/api/portraits/women/85.jpg"
					title="Sandra Adams"
					subtitle="sandra_a88@gmailcom"
				></v-list-item>
			</v-list>

			<v-divider></v-divider>

			<v-list density="compact" nav>
				<v-list-item prepend-icon="mdi-folder" title="My Files" value="myfiles"></v-list-item>
				<v-list-item prepend-icon="mdi-account-multiple" title="Shared with me" value="shared"></v-list-item>
				<v-list-item prepend-icon="mdi-star" title="Starred" value="starred"></v-list-item>
			</v-list>
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
