// router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import LogIn from '../components/LogIn.vue';
import SignUp from '../components/SignUp.vue';
import DashBoard from '../components/DashBoard.vue';
import ReserVation from '../components/ReserVation.vue';
import AccountManagement from '../components/AccountManagement.vue';
import PayMent from '../components/PayMent.vue'
import FaceRecognition from '../components/FaceRecognition.vue'
import NavigationManager from   '../components/NavigationManager.vue'
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
    name: 'DashBoard',
    component: DashBoard
  },
  {
    path: '/account-management',
    name: 'AccountManagement',
    component: AccountManagement
  },
  {
    path: '/reservation',
    name: 'ReserVation',
    component: ReserVation
  },
  {
    path: '/payment',
    name: 'PayMent',
    component: PayMent
  },
  {
    path: '/face-recognition',
    name: 'FaceRecognition',
    component: FaceRecognition
  },
  {
    path: '/navigation',
    name: 'Navigation',
    component: NavigationManager
  }

  // ...其他路由...
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
