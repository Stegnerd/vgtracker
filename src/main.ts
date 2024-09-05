import Lara from "@primevue/themes/lara";
import PrimeVue from "primevue/config";
import "virtual:uno.css";
import { createApp } from "vue";
import App from "./App.vue";
import router from "./routes";

const app = createApp(App);
app.use(PrimeVue, {
  theme: {
    preset: Lara,
    options: {
      darkModeSelector: ".app-dark"
    }
  }
});
app.use(router);

app.mount("#app");
