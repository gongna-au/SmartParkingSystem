<template>
    <div class="sidebar">
      <ul>
        <li><router-link to="/reservation"><v-icon>mdi-parking</v-icon> 车位预订</router-link></li>
        <li><router-link to="/navigation"><v-icon>mdi-map-marker-path</v-icon> 寻车导航</router-link></li>
        <li><router-link to="/payment"><v-icon>mdi-wallet</v-icon> 支付宝支付</router-link></li>
        <li><router-link to="/face-recognition"><v-icon>mdi-face-recognition</v-icon> 人脸识别</router-link></li>
        <li><router-link to="/account-management"><v-icon>mdi-account-cog</v-icon> 账户管理</router-link></li>
      </ul>
    </div>
  </template>
  

<script>
export default {
    name: 'SideBar'
}
</script>

<style scoped>
.sidebar {
    background-color: #34495e; /* 更深的蓝色背景 */
    color: #eaeaea;
    min-width: 250px;
    height: 100vh;
    padding: 20px;
}

.sidebar ul {
    list-style-type: none;
    padding: 0;
}

.sidebar li {
    margin: 15px 0;
    padding: 10px;
    border-radius: 5px;
    transition: background-color 0.3s, color 0.3s;
    display: flex;
    align-items: center;
}

.sidebar li i {
    margin-right: 10px; /* 图标和文本之间的距离 */
}

.sidebar li:hover {
    background-color: #4a6572; /* 深灰蓝色 */
    color: #ffffff;
}

.router-link-exact-active {
    background-color: #1abc9c; /* 激活时的背景颜色 */
    color: #ffffff;
}


</style>
