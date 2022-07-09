<template>
  <div class="home">
    <aside class="show-menu menu">
      <ul class="menu_list menu-list">
        <li class="menu_item">
          <div class="menu_icon">
            <img src="../../assets/img/home_icon/duxiaoshuo.png" alt="">
          </div>
          <div class="menu_kp_link">
            <router-link class="menu-link" to="/home/">起点月票榜</router-link>
          </div>
        </li>

        <li class="menu_item">
          <div class="menu_icon">
            <img src="../../assets/img/home_icon/duxiaoshuo.png" alt="文本编辑器">
          </div>
          <div class="menu_kp_link">
            <router-link class="menu-link" to="/note_edit">文本编辑器</router-link>
          </div>
        </li>

        <li class="menu_item">
          <div class="menu_icon">
            <img src="../../assets/img/home_icon/duxiaoshuo.png" alt="读写">
          </div>
          <div class="menu_kp_link">
            <router-link class="menu-link" to="/read_write">读写操作</router-link>
          </div>
        </li>

        <li>
          <h2>测试</h2>
          <ul>
            <li class="menu_item">
              <div class="menu_icon">
                <img src="../../assets/img/home_icon/box.png" alt="盒子">
              </div>
              <div class="menu_kp_link">
                <router-link class="menu-link" to="/class_box">分类盒子</router-link>
              </div>
            </li>

            <li class="menu_item">
              <div class="menu_icon">
                <img src="../../assets/img/home_icon/mobankuangjia.png" alt="模板">
              </div>
              <div class="menu_kp_link">
                <router-link class="menu-link" to="/class_stl">分类模板</router-link>
              </div>
            </li>

            <li class="menu_item">
              <div class="menu_icon">
                <img src="../../assets/img/home_icon/daochu.png" alt="分类导出">
              </div>
              <div class="menu_kp_link">
                <router-link class="menu-link" to="/class_export">分类导出</router-link>
              </div>
            </li>

            <li class="menu_item">
              <div class="menu_icon">
                <img src="../../assets/img/home_icon/daoru.png" alt="分类导入">
              </div>
              <div class="menu_kp_link">
                <router-link class="menu-link" to="/class_import">分类导入</router-link>
              </div>
            </li>

            <li class="menu_item">
              <div class="menu_icon">
                <img src="../../assets/img/home_icon/duxiaoshuo.png" alt="日记功能">
              </div>
              <div class="menu_kp_link">
                <router-link class="menu-link" to="/note_edit">日记功能</router-link>
              </div>
            </li>

            <li class="menu_item">
              <div class="menu_icon">
                <img src="../../assets/img/home_icon/QQ.png" alt="qq群助手">
              </div>
              <div class="menu_kp_link">
                <router-link class="menu-link" to="/qq_assistant">qq群助手</router-link>
              </div>
            </li>

            <li class="menu_item">
              <div class="menu_icon">
                <img src="../../assets/img/home_icon/bianjiqi.png" alt="拟态UI">
              </div>
              <div class="menu_kp_link">
                <router-link class="menu-link" to="/mimicry_menu">拟态UI</router-link>
              </div>
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
import $ from 'jquery'

export default {
  name: 'Home',
  methods: {
    init() {
      // 获取所有class为menu-获取所有class为menu-link的元素
      let menu_links = $('.menu_list .menu-link')
      menu_links.eq(0).addClass('is-active')
      $('.menu_icon').eq(0).addClass('menu_icon_select')
      // 遍历所有menu-link元素
      menu_links.click(function () {
        // 移除所有class为menu-link的元素的class为active
        let item_link = $(this).parent().prev()
        $('.menu-link.is-active').removeClass('is-active')
        $('.menu_icon').removeClass('menu_icon_select')
        // 给当前点击的元素添加class为active
        this.classList.add('is-active')
        // 找到该元素父亲的上一个兄弟
        item_link.addClass('menu_icon_select')
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
  min-height: 100vh;
  height: 100%;
  padding-top: @Menu_H * 1.5;
  position: relative;


  .show-menu {
    padding: 0 10px;
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

    .menu_item {
      display: flex;

      .menu_icon {

        padding: 0px 5px 0px 5px;
        border-radius: 4px;
        line-height: 40px;
        margin-right: 10px;
        cursor: pointer;

        img {
          width: 20px;
          height: auto;
        }

        &:hover {
          background-color: rgba(0, 0, 0, 0.2);
        }

      }

      // 选中menu_icon时触发的标签
      .menu_icon_select {
        border-left: 2px solid #91d5ff;
        background-color: rgba(0, 0, 0, 0.2);
      }

      .menu_kp_link {
        flex: 0.8;
      }
    }
  }

  // 当界面小于某个值时，菜单收起来
  @media (max-width: 896px) {
    .menu_kp_link {
      display: none;
    }
  }
}
</style>