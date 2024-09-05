import Lara from "@primevue/themes/lara";
import "virtual:uno.css";
//import Aura from "@primevue/themes/aura";
import Button from "primevue/button";
import PrimeVue from "primevue/config";
import { createApp } from "vue";
import App from "./App.vue";

const app = createApp(App);
app.component("Button", Button);
app.use(PrimeVue, {
  theme: {
    preset: Lara,
    options: {
      darkModeSelector: ".app-dark"
    }
  }
});

app.mount("#app");
