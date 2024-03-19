<script setup lang="ts">
import {useConfigStore} from "@/stores/configStore.ts";
import {Theme} from '../../src-tauri/bindings/config/theme.ts'
import * as yup from 'yup';
import {useForm} from "vee-validate";

const store = useConfigStore();
const s = store.configuration;
console.warn('config from store:', s);

// https://vee-validate.logaretm.com/v4/examples/ui-libraries/
const schema = yup.object({
  twitchClientId: yup.string().required(),
  twitchClientSecret: yup.string().required(),
  theme: yup.string<Theme>().required()
})

const {defineField, handleSubmit, errors} = useForm({
  validationSchema: schema
})

const [twitchClientId] = defineField('twitchClientId')
// const [twitchClientSecret] = defineField('twitchClientSecret')
// const [theme] = defineField('theme')

const onSubmit = handleSubmit((values) => {
  console.warn('submitted values', values)
})

</script>

<template>
  <form @submit="onSubmit">
    <div class="flex flex-col gap-2">
      <label for="twitchClientId">twitchClientId</label>
      <InputText id="twitchClientId" v-model="twitchClientId" aria-describedby="twitchClientId-help"/>
      <small id="twitchClientId-help">{{ errors.twitchClientId }}</small>
    </div>
  </form>
</template>
