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

  const toast = useToast();

  const schema = yup.object({
    twitchClientID: yup.string().required().min(1).label("Full name"),
    twitchClientSecret: yup.string().required().min(1).label("Password")
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
  <div class="h-full w-full flex flex-col card">
    <form class="flex flex-col gap-4" @submit="onSubmit">
      <FloatLabel variant="on">
        <InputText id="twitchClientID" v-model="twitchClientID" class="w-lg" />
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
        <Button :disabled="!meta.valid" label="Save" type="submit" />
        <p>{{ meta.valid }}</p>
      </div>
    </form>
  </div>
</template>

<style></style>
