<template>
  <div id="app">
    <nav class="navbar">
      <div class="logo">智慧停车系统</div>
      <div class="nav-links">
        <router-link to="/login" class="nav-item">登录</router-link>
        <router-link to="/signup" class="nav-item">注册</router-link>
        <!-- 更多导航项 -->
      </div>
    </nav>
    <router-view></router-view>
  </div>
</template>

<script>
// 无需特定脚本
</script>

<style>
body, html {
  margin: 0;
  padding: 0;
  height: 100%;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

#app {
  height: 100%;
  background: url('./images/index.png') no-repeat center center fixed; 
  background-size: cover;
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.6); /* 半透明背景 */
  color: white;
  padding: 10px 20px;
}

.logo {
  font-size: 1.5em;
  font-weight: bold;
}

.nav-links {
  display: flex;
}

.nav-item {
  margin: 0 15px;
  color: white;
  text-decoration: none;
  font-size: 1em;
  transition: color 0.3s;
}

.nav-item:hover {
  color: #4CAF50;
}

.router-link-active {
  color: #4CAF50;
}
</style>
