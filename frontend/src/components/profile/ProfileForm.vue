<template>
	<v-form ref="form" v-model="formValid" lazy-validation class="profile-form">
		<v-row>
			<v-col cols="10" offset="1">
				<v-text-field v-model="twitchKey" variant="outlined" label="Twitch Key" required></v-text-field>
				<v-text-field v-model="psnConnection" variant="outlined" label="PSN Connection" required></v-text-field>
			</v-col>
		</v-row>
		<v-row>
			<v-col cols="10">
				<v-btn>SAVE</v-btn>
			</v-col>
		</v-row>
	</v-form>
</template>

<script lang="ts" setup>
	import { ref } from 'vue';

	const formValid = ref(null);
	const twitchKey = ref('');
	const psnConnection = ref('');
</script>

<style scoped>
	.profile-form {
		padding-top: 2em;
	}
</style>
