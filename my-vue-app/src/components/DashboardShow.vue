<template>
    <div class="dashboard">
      <h1>智慧停车系统</h1>
      <div class="dashboard-content">
        <section class="parking-status">
          <h2>车位状态</h2>
          <!-- 车位状态相关内容 -->
        </section>
        <section class="user-info">
          <h2>用户信息</h2>
          <!-- 用户信息相关内容 -->
        </section>
        <section class="history">
          <h2>历史记录</h2>
          <!-- 停车历史记录相关内容 -->
        </section>
        <!-- 可以根据需要添加更多功能区域 -->
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
    return {
      
    };
  },
  }
  </script>
  
  <style>
  .dashboard {
    padding: 20px;
    text-align: center;
  }
  
  .dashboard-content {
    display: flex;
    justify-content: space-around;
    margin-top: 20px;
  }
  
  section {
    flex-basis: 30%;
    padding: 15px;
    border: 1px solid #ddd;
    border-radius: 10px;
    background-color: white;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
  }
  
  h2 {
    color: #333;
  }
  </style>
  