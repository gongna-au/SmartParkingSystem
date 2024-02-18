// router/index.js
import { createRouter, createWebHistory } from 'vue-router';

import LogIn from '../components/LogIn.vue';
import PayMent from '../components/PayMent.vue'
import SignUp from '../components/SignUp.vue';
import DashBoard from '../components/DashBoard.vue';
import ReserVation from '../components/ReserVation.vue';
import AccountManagement from '../components/AccountManagement.vue';
import FaceRecognition from '../components/FaceRecognition.vue'
import NavigationManager from   '../components/NavigationManager.vue'
import UserCenter from '../components/UserCenter.vue'; // 添加这行导入用户中心组件

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
  },
  // 添加用户中心路由配置
  {
    path: '/user-center',
    name: 'UserCenter',
    component: UserCenter
  }

];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
