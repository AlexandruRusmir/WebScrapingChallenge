import { createApp } from 'vue'
import App from './App.vue'
import 'bootstrap/dist/css/bootstrap.css';
import { createTheme } from '@mui/material/styles';

const theme = createTheme();

const app = createApp(App);
app.provide('theme', theme);
app.mount('#app');
