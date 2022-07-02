<!-- 首页显示菜单 -->
<template>
  <div class="index-root">

    <div class="bg-particle">
      <vue-particles color="#fff"></vue-particles>
    </div>

    <!-- 当前系统用户的头像 -->
    <div class="index-user-head">
      <user-head-portrait-vue></user-head-portrait-vue>
    </div>

    <div class="index-data">
      <div :class="card_load" v-for="(item, index) in cards" :key="index">
        <box-card class="data-card" :card="item"></box-card>
      </div>
    </div>
  </div>
</template>

<script>
import BoxCard from '../../components/card/BoxCard.vue'
import UserHeadPortraitVue from '../../components/user_card/UserHeadPortrait.vue'
export default {
  name: 'Index',
  data() {
    return {
      card_load: ['animate__animated', 'animate__lightSpeedInRight'],
      cards: [
        {
          title: '设定集！',
          describe: "在这里，描绘那个精彩的世界...",
          path: '/',
          image: '/img/leimu2.jpeg',
          is_default: true,
        },
        {
          title: '规则书！',
          describe: "你也不想你的规则没人看吧？",
          path: '/',
          image: '/img/leimu.jpeg',
          is_default: false,
        },
        {
          title: '工具集合！',
          describe: "包含了地图，分布式开团工具哦！",
          path: '/home',
          image: '/img/baokemen.jpg',
          is_default: false,
        },
      ],
      is_bg: true
    }
  },
  components: {
    BoxCard,
    UserHeadPortraitVue,
  },
  created() {

  }
}

</script>

<style lang='less' scoped>
.bg-particle {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 90%;
  z-index: 0;
}

.index-root {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: url('/img/bg1.jpg') no-repeat center center;
  background-size: cover;

  &>div {
    margin-top: 5vh;
  }


  .index-data {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-items: center;
    flex-direction: row;

    .data-card {
      margin-right: 1vw;
      margin-top: 2vw;
    }

  }
}

// 当当前页面大小大于等于1024px时，设置为100vw，否则设置为30vw
@media (min-width: 580px) {

  .index-root {
    // 子元素垂直居中
    justify-content: center;
    height: 100%;
    min-height: 100vh;
  }
}

@media (max-width: 580px) {
  .index-root>div:first-child {
    margin-top: 12vh;
  }
}
</style>
