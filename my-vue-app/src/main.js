import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import VueSweetalert2 from 'vue-sweetalert2';
import 'sweetalert2/dist/sweetalert2.min.css'; // 引入 Sweetalert2 样式

const app = createApp(App);

// 使用 VueSweetalert2
app.use(VueSweetalert2);
app.use(router)
// 挂载 Vue 应用
app.mount('#app');
