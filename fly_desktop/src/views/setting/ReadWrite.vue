<!-- 读写文件的视图层 -->
<template>
    <div>
        <button class="button" @click="read">读文件</button><br><br>
        <button class="button is-dark">写文件</button>
        <div class="view">
            {{ text }}
        </div>
    </div>
</template>

<script>
import readme from '../../dao/readme.js';
export default {
    data() {
        return {
            text: ''
        }
    },
    methods: {
        read() {
            this.text = readme()
        }
    }
}

</script>

<style lang='less' scoped>
</style>
