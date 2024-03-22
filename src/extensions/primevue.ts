import {App} from "vue";
import BadgeDirective from "primevue/badgedirective";
import Ripple from "primevue/ripple";
import Badge from "primevue/badge";
import Menu from "primevue/menu";
import Panel from "primevue/panel";
import PrimeVue from "primevue/config";
import Lara from "../presets/lara"
import InputText from "primevue/inputtext";
import Dropdown from "primevue/dropdown";
import Button from "primevue/button";
import FloatLabel from "primevue/floatlabel";


export function addPrimeVue(app: App<Element>) {
    app.directive("badge", BadgeDirective);
    app.directive("ripple", Ripple);

    // components
    app.component("Badge", Badge);
    app.component('Button', Button)
    app.component("Dropdown", Dropdown);
    app.component('FloatLabel', FloatLabel)
    app.component('InputText', InputText)
    app.component("Menu", Menu);
    app.component("Panel", Panel);

    app.use(PrimeVue, {
        unstyled: true,
        pt: Lara,
        ripple: true
    });
}