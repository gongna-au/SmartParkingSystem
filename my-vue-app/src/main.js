import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store'; // 引入store
import VueSweetalert2 from 'vue-sweetalert2';
import 'sweetalert2/dist/sweetalert2.min.css'; // 引入 Sweetalert2 样式
import { BootstrapVue3 } from 'bootstrap-vue-3'
// Optionally import Bootstrap and BootstrapVue CSS
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue-3/dist/bootstrap-vue-3.css'

import vuetify from './plugins/vuetify'
const app = createApp(App);

app.use(vuetify)
// 使用 VueSweetalert2
app.use(VueSweetalert2);

app.use(BootstrapVue3)

// 使用Vue Router
app.use(router)

// 使用Vuex
app.use(store);

// 挂载 Vue 应用
app.mount('#app');
