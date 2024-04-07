<script setup lang="ts">
import { storeToRefs } from "pinia";
import { useForm } from "vee-validate";
import * as yup from 'yup';

import { useConfigStore, } from "@/stores/configStore.ts";

import { UpdateConfigInput } from "../../src-tauri/bindings/config/UpdateConfigInput.ts";

const store = useConfigStore();
const { configuration } = storeToRefs(store);

// https://vee-validate.logaretm.com/v4/examples/ui-libraries/
const schema = yup.object({
  twitchClientId: yup.string().required(),
  twitchClientSecret: yup.string().required(),
});

const { defineField, errors, } = useForm({
  initialValues: {
    twitchClientId: configuration.value?.twitchClientId,
    twitchClientSecret: configuration.value?.twitchClientSecret,
  },
  validationSchema: schema
});

const [twitchClientId] = defineField('twitchClientId');
const [twitchClientSecret] = defineField('twitchClientSecret');

async function onSubmit() {
  const input = {
    twitchClientId: twitchClientId.value,
    twitchClientSecret: twitchClientSecret.value,
  } as UpdateConfigInput;


  await store.updateConfig(input).then(() => {
    console.warn('component await');
  });
}

</script>

<template>
  <form @submit="onSubmit">
    <div class="flex flex-col">
      <div class="flex flex-col gap-2  pt-6 w-full">
        <FloatLabel>
          <InputText
            id="twitchClientId"
            v-model="twitchClientId"
            aria-describedby="twitchClientId-help"
            class="w-1/2"
          />
          <label for="twitchClientId">Twitch Client ID</label>
        </FloatLabel>
        <small id="twitchClientId-help">{{ errors.twitchClientId }}</small>
      </div>
      <div class="flex flex-col gap-2  pt-6 w-full">
        <FloatLabel>
          <InputText
            id="twitchClientSecret"
            v-model="twitchClientSecret"
            aria-describedby="twitchClientSecret-help"
            class="w-1/2"
          />
          <label for="twitchClientSecret">Twitch Client Secret</label>
        </FloatLabel>
        <small id="twitchClientSecret-help">{{ errors.twitchClientSecret }}</small>
      </div>
    </div>

    <div class="pt-12">
      <Button
        label="Submit"
        type="submit"
      />
    </div>
  </form>
</template>
