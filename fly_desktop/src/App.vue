<template>
  <div id="app">
    <router-view></router-view>
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
  height: 100vh;
  backdrop-filter: blur(50px); 
  background-color: @Global_BG_COLOR;
}
</style>
