<!-- 首页显示菜单 -->
<template>
  <div class="index-root">

    <div class="index-data">
      <div v-for="(item, index) in 3" :key="index">
        <box-card :card="card"></box-card>
      </div>
    </div>

  </div>
</template>

<script>
import BoxCard from '../../components/card/BoxCard.vue'
export default {
  name: 'Index',
  data() {
    return {
      card: {}
    }
  },
  components: {
    BoxCard,
  }
}

</script>

<style lang='less' scoped>
.index-root {

  .index-data {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-items: center;
    flex-direction: row;
  }
}

// 当当前页面大小大于等于1024px时，设置为100vw，否则设置为30vw
@media (min-width: 580px) {

  .index-root {
    // 子元素垂直居中
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }
}
</style>
