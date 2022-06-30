<!-- 设计模板 -->
<template>
  <div class="home-test-root">
    <!-- 首先有导航栏，跟内容栏 -->
    <div class="le-nav">
      <!-- 然后最上面的话可以做成头像栏，菜单栏到设置里面去吧，下面就不做与设置相关的 -->
      <div>模板</div>
      <div>记录</div>
      <div>规则书</div>
      <div>人物卡片制作</div>
      <div>物品卡片制作</div>
    </div>
    <div class="rg-content">
      <!-- 其实我觉得这边可以做成标签页 -->
      <div class="card">

      </div>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      
    }
  },
  methods: {
      getData(){
        
      },
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
    width: 20vw;
    height: 100vh;
    background-color: #ff9c6e;
  }

  .rg-content {
    width: 80vw;
    height: 100vh;
    background-color: aqua;
  }

  .cart {
    border-radius: 5px;
    
  }
}
</style>
