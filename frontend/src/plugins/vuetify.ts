import '@mdi/font/css/materialdesignicons.css';
import 'vuetify/styles';

// Composables
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';

// const myCustomLightTheme = {
// 	dark: false,
// 	colors: {
// 		background: '#FFFFFF',
// 		surface: '#FFFFFF',
// 		primary: '#6200EE',
// 		secondary: '#03DAC6',
// 		error: '#B00020',
// 		info: '#2196F3',
// 		success: '#4CAF50',
// 		warning: '#FB8C00'
// 	}
// };
//
// const myCustomDarkTheme = {
// 	dark: true,
// 	colors: {
// 		background: '#264653',
// 		surface: '#FFFFFF',
// 		primary: '#6200EE',
// 		secondary: '#03DAC6',
// 		error: '#e76f51',
// 		info: '#94d2bd',
// 		success: '#2a9d8f',
// 		warning: '#e9c46a'
// 	}
// };

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
	components,
	directives,
	theme: {
		defaultTheme: 'dark'
	}
});
