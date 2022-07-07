<template>
  <div id="app">
    <div id="app-menu">
      <menu-vue />
    </div>
    <router-view></router-view>

    <!-- 导航栏 -->
    <div class="app-navigation">
      <el-breadcrumb separator="/">

        <el-breadcrumb-item v-for="(item, index) in router_paths" :key="index"
          :to="{ path: item.path, name: item.name }">
          <span class="navbar-text">{{ item.name }}</span>
        </el-breadcrumb-item>

      </el-breadcrumb>
    </div>

  </div>
</template>

<script>
import { mapState } from 'vuex'
import MenuVue from '@/components/menu/Menu.vue'
// 添加点击的文字特效
import { click_special_effects } from '@/comm/specialeffects/specialeffects.js'

export default {
  name: 'app',
  data() {
    return {

    }
  },
  computed: {
    ...mapState(['router_paths'])
  },
  components: {
    MenuVue,

  },
  created() {

  },
  mounted() {
    click_special_effects()
  }
}
</script>

<style lang="less">
#app {}

#app-menu {
  -webkit-app-region: drag;
  background-color: #fff;
  display: flex;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: @Menu_H;
  z-index: 4;
}

.app-navigation {
  position: fixed;
  bottom: 3vh;
  left: 2vw;
  z-index: 999;

  .navbar-text {
    color: @Buttom_Bar_Color;
    cursor: pointer;

    &:hover {
      color: @Buttom_Bar_Color_Hover;
    }
  }
}

@media (max-width: 580px) {
  .app-navigation {
    top: 3vh;
  }
}
</style>
