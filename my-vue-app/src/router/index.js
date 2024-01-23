// router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import LogIn from '../components/LogIn.vue';
import SignUp from '../components/SignUp.vue';
import DashboardShow from '../components/DashboardShow.vue';
const routes = [
  {
    path: '/login',
    name: 'LogIn',
    component: LogIn
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: SignUp
  },
  {
    path: '/dashboard',
    name: 'DashboardShow',
    component: DashboardShow
  },
  // ...其他路由...
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
