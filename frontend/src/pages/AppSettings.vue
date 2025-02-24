<script setup lang="ts">
  import { useToast } from "primevue";
import { useForm } from "vee-validate";
import { onMounted } from "vue";
import * as yup from "yup";
import {
  GetConfig,
  UpdateConfig
} from "../../wailsjs/go/controllers/ConfigController";
import { controllers } from "../../wailsjs/go/models";
import AppCard from "../components/AppCard.vue";

  const toast = useToast();

  const schema = yup.object({
    twitchClientID: yup.string().required().min(1).label("Client ID"),
    twitchClientSecret: yup.string().required().min(1).label("Client Secret")
  });

  const { defineField, handleSubmit, errors, meta } = useForm({
    validationSchema: schema
  });

  const [twitchClientID] = defineField("twitchClientID");
  const [twitchClientSecret] = defineField("twitchClientSecret");

  onMounted(async () => {
    GetConfig().then((result) => {
      twitchClientID.value = result.twitch.clientID;
      twitchClientSecret.value = result.twitch.clientSecret;
    });

    // Search("Pokemon Red").then((out) => {
    //   console.warn('search output', out)
    // })
  });

  const onSubmit = handleSubmit(() => {
    const input = {
      twitch: {
        clientID: twitchClientID.value,
        clientSecret: twitchClientSecret.value
      }
    } as controllers.UpdateConfigInput;

    UpdateConfig(input).then(() => {
      toast.add({
        severity: "success",
        summary: "Config Saved",
        group: "br",
        life: 2000
      });
    });
  });
</script>

<template>
  <AppCard>
    <form
      class="flex flex-col gap-4"
      @submit="onSubmit"
    >
      <FloatLabel variant="on">
        <InputText
          id="twitchClientID"
          v-model="twitchClientID"
          data-test="twitchClientID"
          class="w-lg"
        />
        <label for="twitchClientID">Twitch Client ID</label>
        <Message
          v-if="errors.twitchClientID"
          severity="error"
          size="small"
          variant="simple"
        >
          {{ errors.twitchClientID }}
        </Message>
      </FloatLabel>

      <FloatLabel variant="on">
        <InputText
          id="twitchClientSecret"
          v-model="twitchClientSecret"
          data-test="twitchClientSecret"
          class="w-lg"
        />
        <Message
          v-if="errors.twitchClientSecret"
          severity="error"
          size="small"
          variant="simple"
        >
          {{ errors.twitchClientSecret }}
        </Message>
        <label for="twitchClientSecret">Twitch Client Secret</label>
      </FloatLabel>

      <div class="pt-2.5">
        <Button
          :disabled="!meta.valid"
          label="Save"
          type="submit"
          data-test="submit-button"
        />
        <p>{{ meta.valid }}</p>
      </div>
    </form>
  </AppCard>
</template>

<style></style>
