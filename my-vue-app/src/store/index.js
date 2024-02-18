import { createStore } from 'vuex';

export default createStore({
  state() {
    return {
      isLoggedIn: false,
    };
  },
  mutations: {
    setLoggedIn(state, status) {
      state.isLoggedIn = status;
    },
     // 添加退出登录的mutation
    logout(state) {
      state.isLoggedIn = false;
    },
  },
});

