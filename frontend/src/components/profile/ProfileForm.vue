<template>
	<v-form ref="form" class="profile-form">
		<v-row>
			<v-col cols="10" offset="1">
				<v-text-field v-model="profile.twitchKey" variant="outlined" label="Twitch Key" required></v-text-field>
				<v-text-field v-model="profile.psnNpsso" variant="outlined" label="PSN Connection" required></v-text-field>
			</v-col>
		</v-row>
	</v-form>
</template>

<script lang="ts" setup>
	import { onBeforeMount } from 'vue';
	import { useProfileStore } from '../../stores/profile';
	import { storeToRefs } from 'pinia';

	const store = useProfileStore();

	const { readProfile } = useProfileStore();
	onBeforeMount(() => {
		readProfile();
	});

	const { profile } = storeToRefs(store);
</script>

<style scoped>
	.profile-form {
		padding-top: 2em;
	}
</style>
