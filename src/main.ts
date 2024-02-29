import { createApp } from "vue";
import PrimeVue from "primevue/config";
import "./styles.css";
import App from "./App.vue";
import Panel from "primevue/panel";
import Lara from "./presets/lara";
import router from "./routes";
import Menu from "primevue/menu";
import Badge from "primevue/badge";
import BadgeDirective from "primevue/badgedirective";
import Ripple from "primevue/ripple";
import "primeicons/primeicons.css";

const app = createApp(App);

app.directive("badge", BadgeDirective);
app.directive("ripple", Ripple);

app.component("Badge", Badge);
app.component("Menu", Menu);
app.component("Panel", Panel);

app.use(PrimeVue, {
  unstyled: true,
  pt: Lara
});
app.use(router);
app.mount("#app");
