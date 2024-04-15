<script setup lang="ts">
  import { storeToRefs } from "pinia";
  import { useForm } from "vee-validate";
  import * as yup from "yup";

  import { useConfigStore } from "@/stores/configStore.ts";

  import { UpdateConfigInput } from "../../src-tauri/bindings/config/UpdateConfigInput.ts";

  const store = useConfigStore();
  const { configuration } = storeToRefs(store);
  // https://vee-validate.logaretm.com/v4/examples/ui-libraries/
  const schema = yup.object({
    twitchClientId: yup.string().required(),
    twitchClientSecret: yup.string().required()
  });

  const { defineField, errors, meta } = useForm({
    initialValues: {
      twitchClientId: configuration.value?.twitchClientId,
      twitchClientSecret: configuration.value?.twitchClientSecret
    },
    validationSchema: schema
  });

  const [twitchClientId] = defineField("twitchClientId");
  const [twitchClientSecret] = defineField("twitchClientSecret");

  async function onSubmit() {
    const input = {
      twitchClientId: twitchClientId.value,
      twitchClientSecret: twitchClientSecret.value
    } as UpdateConfigInput;

    await store.updateConfig(input).then(() => {
      console.warn("component await");
    });
  }

  defineExpose({ twitchClientId, twitchClientSecret, errors });
</script>

<template>
  <form @submit="onSubmit">
    <div class="flex flex-col">
      <div class="flex flex-col gap-2 pt-6 w-full">
        <PrimeFloatLabel>
          <PrimeInputText
            id="twitchClientId"
            v-model="twitchClientId"
            aria-describedby="twitchClientId-help"
            data-testid="twitchClientIDInput"
            class="w-1/2"
          />
          <label for="twitchClientId">Twitch Client ID</label>
        </PrimeFloatLabel>
        <small
          id="twitchClientId-help"
          class="twitch-client-id-error"
          data-testid="twitchClientId-error"
        >
          {{ errors.twitchClientId }}
        </small>
      </div>
      <div class="flex flex-col gap-2 pt-6 w-full">
        <PrimeFloatLabel>
          <PrimeInputText
            id="twitchClientSecret"
            v-model="twitchClientSecret"
            aria-describedby="twitchClientSecret-help"
            data-testid="twitchClientSecretInput"
            class="w-1/2"
          />
          <label for="twitchClientSecret">Twitch Client Secret</label>
        </PrimeFloatLabel>
        <small
          id="twitchClientSecret-help"
          data-testid="twitchClientSecret-error"
          >{{ errors.twitchClientSecret }}</small
        >
      </div>
    </div>

    <div class="pt-12">
      <PrimeButton
        :disabled="!meta.valid"
        label="Submit"
        type="submit"
        data-testid="submit-button"
      />
    </div>
  </form>
</template>
