<!-- 首页显示菜单 -->
<template>
  <div class="index-root">

    <!-- 当前系统用户的头像 -->
    <div class="index-user-head">
      <user-head-portrait-vue></user-head-portrait-vue>
    </div>

    <div class="index-data">
      <div :class="card_load" v-for="(item, index) in 3" :key="index">
        <box-card class="data-card" :card="card"></box-card>
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
        card_load:['animate__animated','animate__lightSpeedInRight'],
        card:{}
    }
  },
  components: {
    BoxCard,
    UserHeadPortraitVue
  },
  created(){
   
  }
}

</script>

<style lang='less' scoped>
.index-root {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0 5vh;

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
    height: 100vh;
  }
}

@media (max-width: 580px) {
   .index-root>div:first-child {
    margin-top: 12vh;
  }
}
</style>
