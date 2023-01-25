<template>
	<v-form ref='form' v-model='formValid' lazy-validation class='profile-form'>
		<v-row>
			<v-col cols='10' offset='1'>
				<v-text-field v-model='twitchKey' variant='outlined' label='Twitch Key' required></v-text-field>
				<v-text-field v-model='psnConnection' variant='outlined' label='PSN Connection' required></v-text-field>
			</v-col>
		</v-row>
		<v-row>
			<v-col cols='10'>
				<v-btn @click='save'>SAVE</v-btn>
			</v-col>
		</v-row>
	</v-form>
</template>

<script lang='ts' setup>
	import { ref } from 'vue';
	import { UpdateProfile } from '../../../wailsjs/go/api/ProfileBackend';
	import { models } from '../../../wailsjs/go/models';
	import UpdateProfileInput = models.UpdateProfileInput;
	import * as Wails from '../../../wailsjs/runtime';

	const formValid = ref(null);
	const twitchKey = ref('');
	const psnConnection = ref('');

	function save() {
		const input = {
			twitchKey: twitchKey.value,
			psnNpsso: psnConnection.value
		} as UpdateProfileInput;

		UpdateProfile(input)
			.then((result) => {
				console.warn('success result', result);
				Wails.LogDebug(`result: ${result}`);
			})
			.catch((err) => {
				console.warn('error', err);
				Wails.LogDebug(`result: ${err}`);
			});
	}
</script>

<style scoped>
	.profile-form {
		padding-top: 2em;
	}
</style>
