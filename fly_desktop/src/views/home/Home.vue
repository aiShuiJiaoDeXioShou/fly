<template>
  <div class="home">
    <aside class="show-menu menu">
      <p class="menu-label">数据查看</p>
      <ul class="menu_list menu-list">
        <li>
          <router-link to="/home/">起点月票榜</router-link>
        </li>
        <li>
          <router-link to="/note_edit">文本编辑器</router-link>
        </li>
      </ul>
    </aside>
    <div class="show">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Home',
  components: {

  }
}
</script>

<style lang="less" scoped>
.home {
  display: flex;
  height: 100vh;
  padding-top: @Menu_H * 1.5;
  position: relative;

  .show-menu {
    flex: 0.2;
  }

  .show {
    flex: 0.8;
  }

  .menu_list {
    overflow: hidden;
    text-overflow: ellipsis;
    // 防止换行
    white-space: nowrap;
  }
}
</style>