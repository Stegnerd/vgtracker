import Lara from "@primeuix/themes/lara";
import { ToastService } from "primevue";
import PrimeVue from "primevue/config";
import DialogService from "primevue/dialogservice";
import "virtual:uno.css";
import { createApp } from "vue";
import App from "./App.vue";
import "./assets/styles.scss";
import router from "./routes";

const app = createApp(App);
app.use(DialogService);
app.use(ToastService);
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
