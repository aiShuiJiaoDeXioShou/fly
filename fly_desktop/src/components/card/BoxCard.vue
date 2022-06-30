<!-- 用于展示首页菜单 -->
<template>
    <div class="card-root">
        你好世界
    </div>
</template>

<script>
export default {
    props:["card"],
    data() {
        return {
        }
    }
}

</script>

<style lang='less' scoped>
.card-root {
    max-width: 30vw;
    min-width: 275px;
    height: 30vh;
    box-shadow: 0 15px 25px rgb(0, 0, 0, 0.1);
    backdrop-filter: blur(30px);
}
</style>
