<script setup lang="ts">
	import { useRouter } from 'vue-router';

	const router = useRouter();

	function route(value: string) {
		console.warn('value', value);
		router.push({ path: `/${value}` });
	}
</script>

<template>
	<v-list>
		<v-list-item prepend-icon="mdi-account" title="Account" value="account" @click="route('profile')"></v-list-item>

		<v-divider></v-divider>

		<v-list-item
			prepend-icon="mdi-controller"
			title="Game List"
			value="gallery"
			@click="route('gallery')"
		></v-list-item>
		<v-list-item
			prepend-icon="mdi-format-list-bulleted-square"
			title="My Games"
			value="myGames"
			@click="route('my-games')"
		></v-list-item>
	</v-list>
</template>

<style scoped></style>
