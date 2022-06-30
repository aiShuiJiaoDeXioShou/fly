<template>
  <div id="app">
    <router-view></router-view>

    <div class="app-navigation">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item to="/login">登录\注册页面</el-breadcrumb-item>
      </el-breadcrumb>
    </div>

  </div>
</template>

<script>

export default {
  name: 'app',
  data() {
    return {

    }
  },
  components: {

  },
  created() {

  },
}
</script>

<style lang="less">
body {
  width: 100vw;
  height: 100%;
  backdrop-filter: blur(50px);
  background-color: @Global_BG_COLOR;
}
.app-navigation {
  position: fixed;
  bottom: 3vh;
  left: 2vw;
}

@media (max-width: 580px) {
  .app-navigation {
    top: 3vh;
  }
}
</style>
