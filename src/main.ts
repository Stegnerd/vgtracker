import Lara from "@primevue/themes/lara";
import PrimeVue from "primevue/config";
import "virtual:uno.css";
import { createApp } from "vue";
import App from "./App.vue";

const app = createApp(App);
app.use(PrimeVue, {
  theme: {
    preset: Lara,
    options: {
      darkModeSelector: ".app-dark"
    }
  }
});

app.mount("#app");
