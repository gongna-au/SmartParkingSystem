<template>
    <div class="sidebar">
      <ul>
        <li><router-link to="/reservation">车位预订</router-link></li>
        <li><router-link to="/navigation">寻车导航</router-link></li>
        <li><router-link to="/payment">支付宝支付</router-link></li>
        <li><router-link to="/face-recognition">人脸识别</router-link></li>
        <li><router-link to="/account-management">账户管理</router-link></li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    name: 'SideBar'
  }
  </script>
  
  <style scoped>
  .sidebar ul {
    list-style-type: none;
    padding: 0;
  }
  .sidebar li {
    margin: 10px 0;
  }
  </style>
  