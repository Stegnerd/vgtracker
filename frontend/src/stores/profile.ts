import { defineStore } from 'pinia';
import { models } from '../../wailsjs/go/models';
import { ReadProfile, UpdateProfile } from '../../wailsjs/go/api/ProfileBackend';
import * as Wails from '../../wailsjs/runtime';
import { ref } from 'vue';
import Profile = models.Profile;
import ReadProfileOutput = models.ReadProfileOutput;
import UpdateProfileInput = models.UpdateProfileInput;

export const useProfileStore = defineStore('profile', () => {
	const profile = ref<Profile>(new Profile());

	const readProfile = async () => {
		console.warn('calling read profile');
		await ReadProfile()
			.then((result: ReadProfileOutput) => {
				console.warn('profile result', result);
				profile.value = result.userProfile;
			})
			.catch((err) => {
				console.warn('error', err);
				Wails.LogDebug(`result: ${err}`);
			});
	};

	const updateProfile = async () => {
		const input = {
			twitchKey: profile.value.twitchKey,
			psnNpsso: profile.value.psnNpsso
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
	};

	return {
		profile,
		readProfile,
		updateProfile
	};
});
