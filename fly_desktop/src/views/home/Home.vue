<template>
  <div class="home">
    <aside class="show-menu menu">
      <ul class="menu_list menu-list">
        <li>
          <router-link class="menu-link" to="/home/">起点月票榜</router-link>
        </li>
        <li>
          <router-link class="menu-link" to="/note_edit">文本编辑器</router-link>
        </li>
        <li>
          <router-link class="menu-link" to="/read_write">读写操作</router-link>
        </li>
        <li>
          <p>分类功能测试用例</p>
          <ul>
            <li>
              <router-link class="menu-link" to="/class_box">分类盒子</router-link>
            </li>
            <li>
              <router-link class="menu-link" to="/class_stl">分类模板</router-link>
            </li>
            <li>
              <router-link class="menu-link" to="/class_export">分类导出</router-link>
            </li>
            <li>
              <router-link class="menu-link" to="/class_import">分类导入</router-link>
            </li>
            <li>
              <router-link class="menu-link" to="/note_edit">日记功能</router-link>
            </li>
            <li>
              <router-link class="menu-link" to="/qq_assistant">qq群助手</router-link>
            </li>
            <li>
              <router-link class="menu-link" to="/mimicry_menu">拟态UI</router-link>
            </li>
          </ul>
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
  methods: {
    init() {
      // 获取所有class为menu-获取所有class为menu-link的元素
      let menu_links = document.querySelectorAll('.menu_list .menu-link')
      menu_links[0].classList.add('is-active')
      // 遍历所有menu-link元素
      menu_links.forEach(item => {
        item.addEventListener('click', () => {
          // 移除所有class为menu-link的元素的class为active
          let active = document.querySelector('.menu-link.is-active')
          if (active)
            active.classList.remove('is-active')
          // 给当前点击的元素添加class为active
          item.classList.add('is-active')
        })
      })
    }
  },
  components: {

  },
  mounted() {
    this.init()
  }
}
</script>

<style lang="less" scoped>
.home {
  display: flex;
  height: 100%;
  padding-top: @Menu_H * 1.5;
  position: relative;


  .show-menu {
    flex: 0.2;
    padding: 0 10px;
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