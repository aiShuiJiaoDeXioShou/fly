<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import HelloWorld from '@/components/HelloWorld.vue'
import axios from 'axios'

export default {
  name: 'Home',
  components: {
    HelloWorld
  },
  created(){
    axios.get('/api/ok').then(res=>{
      console.log(res);
    }).catch(()=>{
      this.$message({
        type:"error",
        message:"请求失败！"
      })
    })
  }
}
</script>
