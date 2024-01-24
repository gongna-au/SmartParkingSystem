<template>
  <div id="dashboard">
    <sidebar></sidebar>
    <div class="main-content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Sidebar from './SideBar'  // 确保路径的大小写与文件名一致

export default {
  name: 'DashBoard',
  components: {
    Sidebar
  }
}
</script>

<style>
#dashboard {
  display: flex;
}

.main-content {
  flex-grow: 1;
}
</style>
