import Lara from "@/presets/lara";
import PrimeVue from "primevue/config";
import { createApp } from "vue";
import App from "./App.vue";

const app = createApp(App);
app.use(PrimeVue, {
  unstyled: true,
  pt: Lara
});

app.mount("#app");
