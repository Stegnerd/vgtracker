import PrimeVue from "primevue/config";
import { App } from "vue";

import Lara from "../presets/lara";

export function addPrimeVue(app: App<Element>) {
  // components are imported via unplugin plugin
  app.use(PrimeVue, { ripple: true, unstyled: true, pt: Lara });
}
