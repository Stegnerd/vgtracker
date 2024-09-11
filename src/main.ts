import Lara from "@primevue/themes/lara";
import PrimeVue from "primevue/config";
import DialogService from "primevue/dialogservice";
import "virtual:uno.css";
import { createApp } from "vue";
import App from "./App.vue";
import "./assets/styles.scss";
import router from "./routes";

const app = createApp(App);
app.use(DialogService);
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
