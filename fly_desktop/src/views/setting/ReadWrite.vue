<!-- 读写文件的视图层 -->
<template>
    <div class="go_through_motion">
        <button class="button" @click="read">读文件</button><br><br>
        <button class="button is-dark">写文件</button>
        <div class="view">
            {{ text }} <br>
            {{ config.user.name }}
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import readme from '../../dao/readme.js';
export default {
    data() {
        return {
            text: '',
            config: {
                user: {}
            }
        }
    },
    methods: {
        read() {
            this.text = readme()
        }
    },
    mounted() {
        // 从后端接收配置文件的信息
        axios.get('http://localhost:41/api/config/init').then(res => {
            // 将res.data由json字符串转为对象
            this.config = JSON.parse(res.data.data)
        })
    }
}

</script>

<style lang='less' scoped>
</style>
