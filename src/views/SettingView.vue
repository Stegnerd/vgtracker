<script setup lang="ts">
import {useConfigStore} from "@/stores/configStore.ts";
import * as yup from 'yup';
import {useForm} from "vee-validate";

const store = useConfigStore();
const cfg = store.configuration;

interface ThemeOption {
  name: string;
  value: string;
}

const themeOptions = ['Dark', 'Light'].map((o) => ({
  name: o,
  value: o
}));

// https://vee-validate.logaretm.com/v4/examples/ui-libraries/
const schema = yup.object({
  twitchClientId: yup.string().required(),
  twitchClientSecret: yup.string().required(),
  theme: yup.string().required()
})
// TODO: https://tailwind.primevue.org/dropdown/#basic here for theme swutcher
const {defineField, handleSubmit, errors,} = useForm({
  initialValues: {
    twitchClientId: cfg?.twitchClientId,
    twitchClientSecret: cfg?.twitchClientSecret,
    theme: themeOptions[themeOptions.findIndex((f: ThemeOption) => f.name === cfg?.theme)].value
  },
  validationSchema: schema
})

const [twitchClientId] = defineField('twitchClientId')
const [twitchClientSecret] = defineField('twitchClientSecret')
const [theme] = defineField('theme')

const onSubmit = handleSubmit((values) => {
  console.warn('submitted values', values)
})

</script>

<template>
  <form @submit="onSubmit">
    <div class="flex flex-col">
      <div class="flex flex-col gap-2  pt-6 w-full">
        <FloatLabel>
          <InputText id="twitchClientId" v-model="twitchClientId" aria-describedby="twitchClientId-help" class="w-1/2"/>
          <label for="twitchClientId">Twitch Client ID</label>
        </FloatLabel>
        <small id="twitchClientId-help">{{ errors.twitchClientId }}</small>
      </div>
      <div class="flex flex-col gap-2  pt-6 w-full">
        <FloatLabel>
          <InputText id="twitchClientSecret" v-model="twitchClientSecret" aria-describedby="twitchClientSecret-help"
                     class="w-1/2"/>
          <label for="twitchClientSecret">Twitch Client Secret</label>
        </FloatLabel>
        <small id="twitchClientSecret-help">{{ errors.twitchClientSecret }}</small>
      </div>
      <div class="flex flex-col gap-2 pt-6 w-1/5">
        <!--
        this is bugged for dropdown right now
        <FloatLabel>
        -->
        <Dropdown input-id="theme" v-model="theme" :options="themeOptions" aria-describedby="theme-help"
                  optionLabel="name" optionValue="value"/>
        <small id="theme-help">{{ errors.theme }}</small>
      </div>
    </div>

    <div class="pt-12">
      <Button label="Submit" type="submit"/>
    </div>
  </form>
</template>
