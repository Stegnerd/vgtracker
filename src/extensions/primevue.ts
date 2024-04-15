import Badge from "primevue/badge";
import BadgeDirective from "primevue/badgedirective";
import Button from "primevue/button";
import PrimeVue from "primevue/config";
import Dropdown from "primevue/dropdown";
import FloatLabel from "primevue/floatlabel";
import InputText from "primevue/inputtext";
import Menu from "primevue/menu";
import Panel from "primevue/panel";
import Ripple from "primevue/ripple";
import { App } from "vue";

import Lara from "../presets/lara";

export function addPrimeVue(app: App<Element>) {
  // components are imported via unplugin plugin
  app.use(PrimeVue, { ripple: true, unstyled: true, pt: Lara });

  app.directive("badge", BadgeDirective);
  app.directive("ripple", Ripple);

  // components
  app.component("PrimeBadge", Badge);
  app.component("PrimeButton", Button);
  app.component("PrimeDropdown", Dropdown);
  app.component("PrimeFloatLabel", FloatLabel);
  app.component("PrimeInputText", InputText);
  app.component("PrimeMenu", Menu);
  app.component("PrimePanel", Panel);
}
