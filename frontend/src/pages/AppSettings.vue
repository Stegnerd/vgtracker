<script setup lang="ts">
import { Form, FormSubmitEvent, } from "@primevue/forms";
import { zodResolver } from "@primevue/forms/resolvers/zod";
import { useToast } from "primevue";
import { onMounted, reactive, ref } from "vue";
import { z } from "zod";
import { GetConfig, UpdateConfig } from "../../wailsjs/go/controllers/ConfigController";
import { controllers } from "../../wailsjs/go/models";
import AppCard from "../components/AppCard.vue";

const toast = useToast();

const formReady = ref(false)
const accordionOpen = ref<number[]>([0,1]);


type FormData = {
  twitchClientID: string;
  twitchClientSecret: string;
  steamID: string;
  steamApiKey: string;
}

const initialValues = reactive<FormData>({
  twitchClientID: '',
  twitchClientSecret: '',
  steamID: '',
  steamApiKey: ''
});


const onFormSubmit = (e: FormSubmitEvent) => {
  const values = e.values as FormData
  if (e.valid) {
    const input = {
      twitch: {
        clientID: values.twitchClientID,
        clientSecret: values.twitchClientSecret
      },
      steam: {
        steamID: values.steamID,
        apiKey: values.steamApiKey
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
  }
}

const expandAll = () => {
  if (accordionOpen.value.length > 0) {
    accordionOpen.value = [];
  }else {
    accordionOpen.value = [0,1];
  }
}

onMounted(async () => {
  GetConfig().then((result) => {
    initialValues.twitchClientID = result.twitch.clientID;
    initialValues.twitchClientSecret = result.twitch.clientSecret;
    initialValues.steamID = result.steam.steamID;
    initialValues.steamApiKey = result.steam.apiKey;

    formReady.value = true;
  });

});

const resolver = zodResolver(
    z.object({
      twitchClientID: z.string().min(1, {message: "Twitch Client ID is required."}),
      twitchClientSecret: z.string().min(1, {message: "Twitch Client Secret is required."}),
      steamID: z.string(),
      steamApiKey: z.string()
    })
)
</script>

<template>
  <AppCard>
    <Form
      v-if="formReady"
      :resolver
      :initial-values
      :validate-on-mount="true"
      :validate-on-value-update="true"
      class="flex flex-col gap-4"
      @submit="onFormSubmit"
    >
      <div class="flex flex-row gap-2">
        <Button
          type="submit"
          severity="secondary"
          label="Submit"
          data-test="submit-button"
        />
        <Button
          label="Expand All"
          @click="expandAll"
        />
      </div>

      <Accordion
        ref="accordion"
        :value="accordionOpen"
        :multiple="true"
      >
        <AccordionPanel :value="0">
          <AccordionHeader>Twitch</AccordionHeader>
          <AccordionContent>
            <span class="flex flex-col gap-4">
              <FormField
                v-slot="$field"
                name="twitchClientID"
              >
                <FloatLabel variant="on">
                  <InputText
                    type="text"
                    class="w-lg"
                    data-test="twitchClientID"
                  />
                  <label for="twitchClientID">Twitch Client ID</label>
                </FloatLabel>
                <Message
                  v-if="$field?.invalid"
                  severity="error"
                  size="small"
                >
                  {{ $field.error?.message }}
                </Message>
              </FormField>
              <FormField
                v-slot="$field"
                name="twitchClientSecret"
              >
                <FloatLabel variant="on">
                  <InputText
                    type="text"
                    class="w-lg"
                    data-test="twitchClientSecret"
                  />

                  <label for="twitchClientID">Twitch Client Secret</label>
                </FloatLabel>
                <Message
                  v-if="$field?.invalid"
                  severity="error"
                  size="small"
                >
                  {{ $field.error?.message }}
                </Message>
              </FormField>
            </span>
          </AccordionContent>
        </AccordionPanel>
        <AccordionPanel :value="1">
          <AccordionHeader>Steam</AccordionHeader>
          <AccordionContent>
            <span class="flex flex-col gap-4">
              <FormField
                name="steamID"
              >
                <FloatLabel variant="on">
                  <InputText
                    type="text"
                    class="w-lg"
                    data-test="steamID"
                  />
                  <label for="steamID">Steam ID</label>
                </FloatLabel>
                <!-- <Message
          v-if="$field?.invalid"
          severity="error"
          size="small"
        >
          {{ $field.error?.message }}
        </Message> -->
              </FormField>
              <FormField
                v-slot="$field"
                name="steamApiKey"
              >
                <FloatLabel variant="on">
                  <InputText
                    type="text"
                    class="w-lg"
                    data-test="steamApiKey"
                  />
                  <label for="steamApiKey">Steam API Key</label>
                </FloatLabel>
                <Message
                  v-if="$field?.invalid"
                  severity="error"
                  size="small"
                >
                  {{ $field.error?.message }}
                </Message>
              </FormField>
            </span>
          </AccordionContent>
        </AccordionPanel>
      </Accordion>
    </Form>
  </AppCard>
</template>

<style></style>
