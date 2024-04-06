import { createPinia } from "pinia";
import { createApp } from "vue";
import "./styles.css";
import App from "./App.vue";
import { addPrimeVue } from "./extensions/primevue.ts";
import router from "./routes";
import "primeicons/primeicons.css";
import { useConfigStore } from "./stores/configStore.ts";

const app = createApp(App);

// setup ui library
addPrimeVue(app);

// setup stores
const pinia = createPinia()
const configStore = useConfigStore(pinia);

configStore.getConfig().then(() => {
    const conf = configStore.configuration;
    const root = document.getElementsByTagName('html')[0];
    if (conf?.theme === 'Dark') {
        root.classList.toggle('dark');
    }

    app.use(pinia)
        .use(router)
        .mount("#app");
});

