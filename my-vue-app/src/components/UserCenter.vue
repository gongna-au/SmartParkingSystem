<template>
    <div class="user-center">
      <div class="profile-card">
        <h2>用户中心</h2>
        <div class="user-info">
          <p><strong>用户名:</strong> John Doe</p>
          <p><strong>邮箱:</strong> john.doe@example.com</p>
          <!-- 更多用户信息 -->
        </div>
        <button @click="logout" class="logout-btn">退出登录</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      logout() {
        this.$store.commit('logout'); // 调用Vuex中的logout mutation
        this.$router.push('/login'); // 重定向用户到登录页面
      },
    },
  };
  </script>
  
  <style scoped>
  .user-center {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f5f5f5;
  }
  
  .profile-card {
    background: white;
    padding: 2rem;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 400px;
    text-align: center;
  }
  
  .user-info p {
    font-size: 1rem;
    color: #333;
    margin: 0.5rem 0;
  }
  
  .logout-btn {
    background-color: #007bff;
    color: white;
    padding: 0.75rem 1.5rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    margin-top: 1rem;
    font-size: 1rem;
    transition: background-color 0.2s;
  }
  
  .logout-btn:hover {
    background-color: #0056b3;
  }
  </style>
  