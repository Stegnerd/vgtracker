<template>
	<ProfileForm></ProfileForm>
	<v-row>
		<v-col cols="10">
			<v-btn @click="updateProfile">SAVE</v-btn>
		</v-col>
	</v-row>
</template>

<script setup lang="ts">
	import ProfileForm from '../components/profile/ProfileForm.vue';
	import { useProfileStore } from '../stores/profile';

	const { updateProfile } = useProfileStore();
</script>

<style scoped></style>
