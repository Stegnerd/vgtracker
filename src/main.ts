import { createApp } from "vue";
import PrimeVue from "primevue/config";
import "./styles.css";
import App from "./App.vue";
import Panel from "primevue/panel";
import Lara from "./presets/lara";

const app = createApp(App);
app.component("Panel", Panel);
app.use(PrimeVue, {
  unstyled: true,
  pt: Lara
});
app.mount("#app");
