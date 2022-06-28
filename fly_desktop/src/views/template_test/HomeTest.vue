<!-- 设计模板 -->
<template>
  <div class="home-test-root">
    <!-- 首先有导航栏，跟内容栏 -->
    <div class="le-nav"></div>
    <div class="rg-contet"></div>
  </div>
</template>

<script>
export default {
  data () {
    return {
    }
  }
}

</script>

<style lang='less' scoped>
.home-test-root {
  display: flex;
  background: url("/img/山谷之上.jpg") no-repeat;
  // 背景图片大小与窗口大小一致
  background-size: cover;


  .le-nav {
    width: 20vh;
    height: 100vh;
  }

  .rg-content {
    width: 80vh;
    height: 100vh;
  }
}
</style>
