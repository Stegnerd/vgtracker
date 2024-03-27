import {createApp} from "vue";
import "./styles.css";
import App from "./App.vue";
import router from "./routes";
import "primeicons/primeicons.css";
import {addPrimeVue} from "./extensions/primevue.ts";
import {createPinia} from "pinia";
import {useConfigStore} from "./stores/configStore.ts";

const app = createApp(App);

// setup ui library
addPrimeVue(app)

// setup stores
const pinia = createPinia()
app.use(pinia)

app.use(router);
app.mount("#app");

// populate data before load
const configStore = useConfigStore();
await configStore.getConfig()