import "primeicons/primeicons.css";
import "./styles.css";

import { createPinia } from "pinia";
import { createApp } from "vue";

import App from "./App.vue";
import { addPrimeVue } from "./extensions/primevue.ts";
import router from "./routes";

const app = createApp(App);

// setup ui library
addPrimeVue(app);

// setup stores
const pinia = createPinia();

app.use(pinia).use(router).mount("#app");
