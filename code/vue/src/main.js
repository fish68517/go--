
import { createApp } from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import App from './App.vue'
import '@/assets/css/plugins.css';
import '@/assets/css/style.css';
import '@/assets/css/responsive.css';
const app = createApp(App)
app.use(ElementPlus)
app.mount('#app')
