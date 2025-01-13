<script setup lang="ts">
  // import { invoke } from "@tauri-apps/api/core";
  import { useToast } from "primevue";
  import { useForm } from "vee-validate";
  import { onMounted } from "vue";
  import * as yup from "yup";
  import { VGConfig } from "../models/api";

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
    const t: VGConfig = {
      twitchClientId: "cats",
      twitchClientSecret: "dogs"
    }; // await invoke("read_config");
    console.warn("result", t);
    twitchClientID.value = t.twitchClientId;
    twitchClientSecret.value = t.twitchClientSecret;
  });

  const onSubmit = handleSubmit((values) => {
    console.warn("Submitted with", values);
    toast.add({
      severity: "success",
      summary: "Config Saved",
      group: "br",
      life: 2000
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
